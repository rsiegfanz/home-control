package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/seeder/pkg/seeds/kafkaSeeds"
	"github.com/rsiegfanz/home-control/backend/seeder/pkg/seeds/postgresSeeds"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
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

	dbConfig, kafkaConfig := loadConfigs()

	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	err = postgresSeeds.SeedRooms(db)
	if err != nil {
		logging.Logger.Fatal("Seed postgress rooms error", zap.Error(err))
	}

	err = postgresSeeds.SeedElectricityMeter(db)
	if err != nil {
		logging.Logger.Fatal("Seed postgress electricity meters error", zap.Error(err))
	}

	err = kafkaSeeds.SeedTopics(kafkaConfig)
	if err != nil {
		logging.Logger.Fatal("Seed Kafka topics error", zap.Error(err))
	}

	logging.Logger.Info("Seeder finished")
}

func loadConfigs() (postgres.Config, kafka.Config) {
	/*config, err := config.LoadConfig[config.ConfigPostgres]()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}*/
	postgresConfig := postgres.Config{}
	postgresConfig.Host = "localhost"
	postgresConfig.Port = 5432
	postgresConfig.DbName = "home_control_db"
	postgresConfig.User = "home_control_user"
	postgresConfig.Password = "home_control_password"

	kafkaConfig := kafka.Config{}
	kafkaConfig.Host = "localhost:9092"

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig
}
