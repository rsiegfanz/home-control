package main

import (
	"log"
	"time"

	"github.com/rsiegfanz/home-control/backend/ingestor/pkg/climateMeasurements"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/redis"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	logPath := "/var/log"
	if !config.IsProd() {
		logPath = "d:\\dev\\docker\\share\\home-control\\promtail"
	}

	if err := logging.InitLogger("info", "ingestor", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Ingestor started")

	dbConfig, kafkaConfig, redisConfig := loadConfigs()

	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	rooms := []models.Room{}
	db.Find(&rooms)

	ingestClimateMeasurements(dbConfig, kafkaConfig, redisConfig)

	logging.Logger.Sugar().Infof("rooms", rooms)
}

func ingestClimateMeasurements(postgresConfig postgres.Config, kafkaConfig kafka.Config, redisConfig redis.Config) {
	ingestor := climateMeasurements.NewIngestor(postgresConfig, kafkaConfig, redisConfig)
	defer ingestor.Close()

	for {
		ingestor.Execute()

		time.Sleep(5 * time.Second)
	}
}

func loadConfigs() (postgres.Config, kafka.Config, redis.Config) {
	if config.IsProd() {
		return loadConfigsProd()
	}
	return loadConfigsDev()
}

func loadConfigsProd() (postgres.Config, kafka.Config, redis.Config) {
	logging.Logger.Info("Loading PROD environment")

	postgresConfig, err := config.LoadConfig[postgres.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading postgres config", zap.Error(err))
	}

	kafkaConfig, err := config.LoadConfig[kafka.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading kafka config", zap.Error(err))
	}

	redisConfig, err := config.LoadConfig[redis.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading redis config", zap.Error(err))
	}

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig, redisConfig
}

func loadConfigsDev() (postgres.Config, kafka.Config, redis.Config) {
	logging.Logger.Warn("Loading DEV environment")

	postgresConfig := postgres.Config{}
	postgresConfig.Host = "localhost"
	postgresConfig.Port = 5432
	postgresConfig.DbName = "home_control_db"
	postgresConfig.User = "home_control_user"
	postgresConfig.Password = "home_control_password"

	kafkaConfig := kafka.Config{}
	kafkaConfig.Host = "localhost:9092"

	redisConfig := redis.Config{}
	redisConfig.Host = "localhost:6379"

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig, redisConfig
}
