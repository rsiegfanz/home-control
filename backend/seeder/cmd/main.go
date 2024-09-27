package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/seeds"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	logPath := "d:\\dev\\docker\\share\\home-control\\promtail"
	if err := logging.InitLogger("info", "seeder", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Seeder started")

	dbConfig := postgres.Config{}
	dbConfig.Host = "localhost"
	dbConfig.Port = 5432
	dbConfig.DbName = "home_control_db"
	dbConfig.User = "home_control_user"
	dbConfig.Password = "home_control_password"

	/*config, err := config.LoadConfig[config.ConfigPostgres]()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}*/

	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	rooms := []models.Room{}
	db.Find(&rooms)

	logging.Logger.Sugar().Infof("rooms", rooms)

	err = seeds.SeedRooms(db)
	if err != nil {
		logging.Logger.Fatal("Seed error", zap.Error(err))
	}
}
