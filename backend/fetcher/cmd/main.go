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
)

func main() {
	logPath := "/var/log"
	if !config.IsProd() {
		logPath = "d:\\dev\\docker\\share\\home-control\\promtail"
	}

	if err := logging.InitLogger("info", "fetcher", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Fetcher started")

	_, kafkaConfig, fetcherConfig := loadConfigs()

	switch fetcherConfig.Modus {
	case "climate_current":
		fetchClimateMeasurementsLatest(kafkaConfig, fetcherConfig)
	case "climate_history":
		fetchClimateMeasurementsAll(kafkaConfig, fetcherConfig)
	default:
		logging.Logger.Error("invalid or missing modus", zap.String("modus", fetcherConfig.Modus))
	}

	logging.Logger.Info("Fetcher stopped")
}

func fetchClimateMeasurementsLatest(kafkaConfig kafka.Config, fetcherConfig configs.FetcherConfig) {
	climateMeasurementsFetcher, err := climateMeasurements.NewFetcher(fetcherConfig, kafkaConfig)
	if err != nil {
		logging.Logger.Fatal("Error instantiating room fetcher", zap.Error(err))
	}
	defer climateMeasurementsFetcher.Close()

	cnt := uint32(10)
	for {
		if climateMeasurementsFetcher.FetchLatest(cnt) {
			cnt = 1
		} else {
			cnt++
			cnt = min(cnt, 1000)
		}

		time.Sleep(30 * time.Second)
	}
}

func fetchClimateMeasurementsAll(kafkaConfig kafka.Config, fetcherConfig configs.FetcherConfig) {
	climateMeasurementsFetcher, err := climateMeasurements.NewFetcher(fetcherConfig, kafkaConfig)
	if err != nil {
		logging.Logger.Fatal("Error instantiating room fetcher", zap.Error(err))
	}
	defer climateMeasurementsFetcher.Close()

	climateMeasurementsFetcher.FetchLatest(99999)
}

// func fetchClimateMeasurementsHistory(kafkaConfig kafka.Config, dbConfig postgres.Config, fetcherConfig configs.FetcherConfig) {
// 	climateMeasurementsFetcher, err := climateMeasurements.NewFetcher(fetcherConfig, kafkaConfig)
// 	if err != nil {
// 		logging.Logger.Fatal("Error instantiating room fetcher", zap.Error(err))
// 	}
// 	defer climateMeasurementsFetcher.Close()

// 	db, err := postgres.InitDB(dbConfig)
// 	if err != nil {
// 		logging.Logger.Fatal("Error opening postgres database", zap.Error(err))
// 	}

// 	rooms := []models.Room{}
// 	db.Find(&rooms)

// 	climateMeasurementsFetcher.FetchHistory(rooms)
// }

func loadConfigs() (postgres.Config, kafka.Config, configs.FetcherConfig) {
	if config.IsProd() {
		return loadConfigsProd()
	}
	return loadConfigsDev()
}

func loadConfigsProd() (postgres.Config, kafka.Config, configs.FetcherConfig) {
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

	fetcherConfig, err := config.LoadConfig[configs.FetcherConfig]()
	if err != nil {
		logging.Logger.Fatal("Error loading config", zap.Error(err))
	}

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig, fetcherConfig
}

func loadConfigsDev() (postgres.Config, kafka.Config, configs.FetcherConfig) {
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

	fetcherConfig, err := config.LoadConfig[configs.FetcherConfig]()
	if err != nil {
		logging.Logger.Fatal("Error loading config", zap.Error(err))
	}

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, kafkaConfig, fetcherConfig
}
