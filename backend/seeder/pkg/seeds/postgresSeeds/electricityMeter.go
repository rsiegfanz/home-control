package postgresSeeds

import (
	"fmt"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedElectricityMeter(db *gorm.DB) error {
	electricityMeters := []models.ElectricityMeter{
		{ExternalId: "1 APA01 1612 1851", Name: "OG"},
		{ExternalId: "1 APA01 1612 1824", Name: "EG"},
		{ExternalId: "1 ISK00 9003 8604", Name: "Wärmepumpe"},
	}

	for _, electricityMeter := range electricityMeters {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(&electricityMeter).Error
		if err != nil {
			return fmt.Errorf("failed to upsert electricityMeter with ExternalId %s: %w", electricityMeter.ExternalId, err)
		}
		fmt.Printf("Upserted electricity meter: %s\n", electricityMeter.ExternalId)
	}

	fmt.Println("Successfully upserted electricity meter table.")
	return nil
}
