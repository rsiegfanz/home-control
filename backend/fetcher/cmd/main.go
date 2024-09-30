package main

import (
	"log"
	"time"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/climateMeasurements"
	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
	// "gorm.io/driver/postgres"
)

func main() {
	// logPath := "d:\\dev\\docker\\share\\home-control\\promtail"
	logPath := "/mnt/d/dev/docker/share"
	if err := logging.InitLogger("info", "fetcher", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	if logging.Logger == nil {
		log.Fatalf("Logger not initialized")
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Fetcher started")

	_, kafkaConfig, fetcherConfig := loadConfigs()
	/*
		db, err := postgres.InitDB(dbConfig)
		if err != nil {
			logging.Logger.Fatal("Error opening database", zap.Error(err))
		}

		dbRooms := []models.Room{}
		db.Find(&dbRooms)
	*/
	// logging.Logger.Sugar().Infof("rooms", rooms)
	//

	fetchClimateMeasurements(kafkaConfig, fetcherConfig)

	logging.Logger.Info("Fetcher stopped")
}

func fetchClimateMeasurements(kafkaConfig kafka.Config, fetcherConfig configs.FetcherConfig) {
	climateMeasurementsFetcher, err := climateMeasurements.NewFetcher(fetcherConfig, kafkaConfig)
	if err != nil {
		logging.Logger.Fatal("Error instantiating room fetcher", zap.Error(err))
	}
	defer climateMeasurementsFetcher.Close()

	for {
		climateMeasurementsFetcher.FetchLatest()

		time.Sleep(30 * time.Second)
	}
}

func loadConfigs() (postgres.Config, kafka.Config, configs.FetcherConfig) {
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

	// redisConfig := redis.Config{}
	// redisConfig.Host = "localhost:6379"

	fetcherConfig, err := config.LoadConfig[configs.FetcherConfig]()
	if err != nil {
		logging.Logger.Fatal("Error loading config", zap.Error(err))
	}

	log.Printf("LOGGING CONFIG: %v", fetcherConfig)

	//	fetcherConfig := configs.FetcherConfig{}
	//	fetcherConfig.Url = ""
	//

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig, fetcherConfig
}
