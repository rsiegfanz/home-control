package postgres

import (
	"fmt"
	"log"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.ConfigPostgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Host, config.User, config.Password, config.DbName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return db, err
	}

	db.AutoMigrate(&models.Room{})

	return db, nil
}
