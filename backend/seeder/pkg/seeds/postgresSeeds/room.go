package postgresSeeds

import (
	"fmt"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedRooms(db *gorm.DB) error {
	rooms := []models.Room{
		{ExternalId: "funkbme280", Name: "Garage"},
		{ExternalId: "eacbfbd86e3e", Name: "EG Wohnzimmer"},
		{ExternalId: "8cce4ef1ddea", Name: "EG Küche"},
		{ExternalId: "4091514f739a", Name: "OG Wohnzimmer"},
		{ExternalId: "4091514ef9c0", Name: "OG Küche"},
		{ExternalId: "e89f6d94fb4", Name: "OG Schlafzimmer"},
		{ExternalId: "acbfbd829b2", Name: "OG Empore"},
		{ExternalId: "e868e758ea4f", Name: "UG Gästezimmer"},
		{ExternalId: "8cce4ef2882a", Name: "UG Fitnessraum"},
	}

	for _, room := range rooms {
		err := db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(&room).Error
		if err != nil {
			return fmt.Errorf("failed to upsert room with ExternalId %s: %w", room.ExternalId, err)
		}
		fmt.Printf("Upserted room: %s\n", room.ExternalId)
	}

	fmt.Println("Successfully upserted rooms table.")
	return nil
}
