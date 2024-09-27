package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/rooms"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
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

	dbConfig, kafkaConfig, fetcherConfig := loadConfigs()

	logging.Logger.Debug("Configs loaded")

	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	dbRooms := []models.Room{}
	db.Find(&dbRooms)

	// logging.Logger.Sugar().Infof("rooms", rooms)

	roomFetcher, err := rooms.NewRoomFetcher(fetcherConfig, kafkaConfig)
	if err != nil {
		logging.Logger.Fatal("Error instantiating room fetcher", zap.Error(err))
	}
	roomFetcher.FetchLatest()

	logging.Logger.Info("Logger stopped")
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
	kafkaConfig.Host = "localhost"

	fetcherConfig, err := config.LoadConfig[configs.FetcherConfig]()
	if err != nil {
		logging.Logger.Fatal("Error loading config", zap.Error(err))
	}

	log.Printf("LOGGING CONFIG:", fetcherConfig)

	//	fetcherConfig := configs.FetcherConfig{}
	//	fetcherConfig.Url = ""

	return postgresConfig, kafkaConfig, fetcherConfig
}
