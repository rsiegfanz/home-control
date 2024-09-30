package climateMeasurements

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	rskafka "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka/dtos"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Ingestor struct {
	kafkaReader *kafka.Reader
	db          *gorm.DB
}

func NewIngestor(postgresConfig postgres.Config, configKafka rskafka.Config) *Ingestor {
	reader := rskafka.InitReader(configKafka, rskafka.TopicClimateMeasurements)

	db, err := postgres.InitDB(postgresConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	return &Ingestor{kafkaReader: reader, db: db}
}

func (i *Ingestor) Close() {
	defer i.kafkaReader.Close()
}

func (i *Ingestor) Execute() {
	for {
		message, err := i.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var dto dtos.ClimateMeasurement
		err = json.Unmarshal(message.Value, &dto)
		if err != nil {
			break
		}

		layout := "02.01.2006 15:04:05"
		parsedTime, err := time.Parse(layout, dto.Timestamp)
		if err != nil {
			logging.Logger.Warn("Invalid timestamp %s / error: %v", zap.String("ts", dto.Timestamp), zap.Error(err))
		} else {
			logging.Logger.Info("OK")
			entity := models.ClimateMeasurement{
				RoomExternalId: strings.Replace(dto.RoomId, ".werte", "", -1),
				Temperature:    dto.Temperature,
				Humidity:       dto.Humidity,
				RecordedAt:     parsedTime,
			}

			i.db.Save(&entity)
		}
		err = i.kafkaReader.CommitMessages(context.Background(), message)
		if err != nil {
			logging.Logger.Warn("Commit message error", zap.Error(err))
		}
	}
}
