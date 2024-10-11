package resolvers

import (
	"context"
	"time"

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
