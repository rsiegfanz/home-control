package climateMeasurements

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	rskafka "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka/dtos"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	rsredis "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/redis"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Ingestor struct {
	kafkaReader *kafka.Reader
	db          *gorm.DB
	redisClient *redis.Client
}

func NewIngestor(postgresConfig postgres.Config, configKafka rskafka.Config, configRedis rsredis.Config) *Ingestor {
	kafkaReader := rskafka.InitReader(configKafka, rskafka.TopicClimateMeasurements)
	redisClient := rsredis.InitClient(configRedis)
	db, err := postgres.InitDB(postgresConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}
	return &Ingestor{kafkaReader: kafkaReader, db: db, redisClient: redisClient}
}

func (i *Ingestor) Close() {
	i.kafkaReader.Close()
	i.redisClient.Close()
}

func (i *Ingestor) Execute() {
	ctx := context.Background()

	for {
		message, err := i.kafkaReader.FetchMessage(ctx)
		if err != nil {
			logging.Logger.Error("Error fetching message from Kafka", zap.Error(err))
			break
		}

		entity, err := i.processMessage(message.Value)
		if err != nil {
			logging.Logger.Error("Error processing message", zap.Error(err))
			continue
		}

		if err := i.saveToPostgres(entity); err != nil {
			logging.Logger.Warn("Failed to insert measurement in Postgres", zap.Error(err))
		}

		if err := i.saveToRedis(ctx, entity); err != nil {
			logging.Logger.Warn("Failed to process measurement in Redis", zap.Error(err))
		}

		if err := i.kafkaReader.CommitMessages(ctx, message); err != nil {
			logging.Logger.Warn("Commit message error", zap.Error(err))
		}
	}
}

func (i *Ingestor) processMessage(messageValue []byte) (*models.ClimateMeasurement, error) {
	var dto dtos.ClimateMeasurement
	if err := json.Unmarshal(messageValue, &dto); err != nil {
		return nil, fmt.Errorf("error unmarshalling message: %w", err)
	}

	parsedTime, err := time.Parse("02.01.2006 15:04:05", dto.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp %s: %w", dto.Timestamp, err)
	}

	return &models.ClimateMeasurement{
		RoomExternalId: strings.Replace(dto.RoomId, ".werte", "", -1),
		Temperature:    dto.Temperature,
		Humidity:       dto.Humidity,
		RecordedAt:     parsedTime,
	}, nil
}

func (i *Ingestor) saveToPostgres(entity *models.ClimateMeasurement) error {
	return i.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "recorded_at"}, {Name: "room_external_id"}},
		DoNothing: true,
	}).Create(entity).Error
}

func (i *Ingestor) saveToRedis(ctx context.Context, entity *models.ClimateMeasurement) error {
	lastMeasurementKey := fmt.Sprintf("last_measurement:%s", entity.RoomExternalId)

	var currentMeasurement models.ClimateMeasurement
	currentData, err := i.redisClient.Get(ctx, lastMeasurementKey).Bytes()
	if err == nil {
		if err := json.Unmarshal(currentData, &currentMeasurement); err == nil {
			if !entity.RecordedAt.After(currentMeasurement.RecordedAt) {
				// "new" kafka-value is older than current redis-value -> ignore
				return nil
			}
		}
	} else if err != redis.Nil {
		return fmt.Errorf("error checking current measurement: %w", err)
	}

	value, err := json.Marshal(entity)
	if err != nil {
		return fmt.Errorf("error marshalling entity: %w", err)
	}

	_, err = i.redisClient.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, lastMeasurementKey, value, 0)

		updateChannel := fmt.Sprintf("updates:%s", entity.RoomExternalId)
		pipe.Publish(ctx, updateChannel, value)

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to update and publish measurement: %w", err)
	}

	return nil
}

func (i *Ingestor) updateLastMeasurement(ctx context.Context, entity *models.ClimateMeasurement, redisValue []byte) error {
	lastMeasurementKey := fmt.Sprintf("last_measurement:%s", entity.RoomExternalId)

	currentLastMeasurement, err := i.redisClient.Get(ctx, lastMeasurementKey).Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("failed to get current last measurement from Redis: %w", err)
	}

	shouldUpdate := true
	if err == nil {
		var currentMeasurement models.ClimateMeasurement
		if err := json.Unmarshal([]byte(currentLastMeasurement), &currentMeasurement); err == nil {
			if currentMeasurement.RecordedAt.After(entity.RecordedAt) {
				shouldUpdate = false
			}
		}
	}

	if shouldUpdate {
		if err := i.redisClient.Set(ctx, lastMeasurementKey, redisValue, 0).Err(); err != nil {
			return fmt.Errorf("failed to update last measurement in Redis: %w", err)
		}
	}

	return nil
}
