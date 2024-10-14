package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
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

	measurements := make(chan interface{})
	go func() {
		defer close(measurements)
		pubsub := r.RedisClient.Subscribe(params.Context, fmt.Sprintf("updates:%s", roomExternalId))
		defer pubsub.Close()

		for {
			select {
			case <-params.Context.Done():
				return
			case msg := <-pubsub.Channel():
				var measurement models.ClimateMeasurement
				if err := json.Unmarshal([]byte(msg.Payload), &measurement); err != nil {
					// Log the error, but continue
					continue
				}
				measurements <- measurement
			}
		}
	}()

	return measurements, nil
}
