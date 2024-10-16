package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func (r *QueryResolver) GetClimateMeasurements(ctx context.Context, startDate string, endDate string, roomExternalId string) ([]*models.ClimateMeasurement, error) {
	var measurements []*models.ClimateMeasurement
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		return nil, err
	}

	err = r.DB.Where("recorded_at BETWEEN ? AND ? AND room_external_id = ?", start, end, roomExternalId).
		Find(&measurements).Error
	if err != nil {
		return nil, err
	}

	return measurements, nil
}

func (r *QueryResolver) SubscribeToClimateMeasurements(params graphql.ResolveParams) (interface{}, error) {
	roomExternalId, ok := params.Args["roomExternalId"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid roomExternalId")
	}

	logging.Logger.Info("Starting subscription for room", zap.String("roomExternalId", roomExternalId))

	return func() (interface{}, error) {
		pubsub := r.RedisClient.Subscribe(params.Context, fmt.Sprintf("updates:%s", roomExternalId))
		defer pubsub.Close()

		msg, err := pubsub.ReceiveMessage(params.Context)
		if err != nil {
			logging.Logger.Error("Error receiving message from Redis", zap.Error(err))
			return nil, err
		}

		var measurement models.ClimateMeasurement
		if err := json.Unmarshal([]byte(msg.Payload), &measurement); err != nil {
			logging.Logger.Error("Error unmarshalling measurement", zap.Error(err))
			return nil, err
		}

		logging.Logger.Debug("Received new measurement from Redis", zap.Any("measurement", measurement))

		return &measurement, nil
	}, nil
}
