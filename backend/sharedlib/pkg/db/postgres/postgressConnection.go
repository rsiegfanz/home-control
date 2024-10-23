package postgres

import (
	"fmt"
	"log"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		config.Host, config.User, config.Password, config.DbName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return db, err
	}

	db.AutoMigrate(&models.Room{}, &models.ElectricityMeter{}, &models.ClimateMeasurement{})

	convertToHypertable(db, "climate_measurements", "recorded_at")

	return db, nil
}

func convertToHypertable(db *gorm.DB, table string, column string) {
	var result int64
	db.Raw("SELECT COUNT(1) FROM timescaledb_information.hypertables WHERE hypertable_name = ?", table).Scan(&result)

	if result == 0 {
		logging.Logger.Info("Converting table %s into hypertable", zap.String("table", table))
		db.Exec("SELECT create_hypertable(?, ?)", table, column)
	}
	logging.Logger.Debug("Table %s already converted into hypertable", zap.String("table", table))
}
