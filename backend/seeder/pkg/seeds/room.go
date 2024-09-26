package seeds

import (
	"fmt"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SeedRooms(db *gorm.DB) error {
	rooms := []models.Room{
		{ExternalId: "funkbme280", Name: "Garage"},
		{ExternalId: "eacbfbd86e3e.werte", Name: "EG Wohnzimmer"},
		{ExternalId: "8cce4ef1ddea.werte", Name: "EG Küche"},
		{ExternalId: "4091514f739a.werte", Name: "OG Wohnzimmer"},
		{ExternalId: "4091514ef9c0.werte", Name: "OG Küche"},
		{ExternalId: "e89f6d94fb4.werte", Name: "OG Schlafzimmer"},
		{ExternalId: "acbfbd829b2.werte", Name: "OG Empore"},
		{ExternalId: "e868e758ea4f.werte", Name: "Keller Gästezimmer"},
		{ExternalId: "8cce4ef2882a.werte", Name: "Keller Fitnessraum"},
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
