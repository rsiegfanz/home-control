package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/seeder/pkg/seeds/kafkaSeeds"
	"github.com/rsiegfanz/home-control/backend/seeder/pkg/seeds/postgresSeeds"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	//logPath := "d:\\dev\\docker\\share\\home-control\\promtail"
	if err := logging.InitLogger("info", "seeder", ""); err != nil {
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
	if config.IsProd() {
		return loadConfigsProd()
	}
	return loadConfigsDev()
}

func loadConfigsProd() (postgres.Config, kafka.Config) {
	logging.Logger.Info("Loading PROD environment")

	postgresConfig, err := config.LoadConfig[postgres.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading postgres config", zap.Error(err))
	}

	kafkaConfig, err := config.LoadConfig[kafka.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading kafka config", zap.Error(err))
	}

	// redisConfig := redis.Config{}
	// redisConfig.Host = "localhost:6379"

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig
}

func loadConfigsDev() (postgres.Config, kafka.Config) {
	logging.Logger.Warn("Loading DEV environment")

	postgresConfig := postgres.Config{}
	postgresConfig.Host = "localhost"
	postgresConfig.Port = 5432
	postgresConfig.DbName = "home_control_db"
	postgresConfig.User = "home_control_user"
	postgresConfig.Password = "home_control_password"

	kafkaConfig := kafka.Config{}
	kafkaConfig.Host = "localhost:9092"

	// redisConfig := redis.Config{}
	// redisConfig.Host = "localhost:6379"

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig
}
