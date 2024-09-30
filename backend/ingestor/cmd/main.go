package main

import (
	"log"
	"time"

	"github.com/rsiegfanz/home-control/backend/ingestor/pkg/climateMeasurements"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/kafka"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	logPath := "d:\\dev\\docker\\share\\home-control\\promtail"
	if err := logging.InitLogger("info", "ingestor", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Ingestor started")

	dbConfig, kafkaConfig := loadConfigs()

	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	rooms := []models.Room{}
	db.Find(&rooms)

	ingestClimateMeasurements(dbConfig, kafkaConfig)

	logging.Logger.Sugar().Infof("rooms", rooms)
}

func ingestClimateMeasurements(postgresConfig postgres.Config, kafkaConfig kafka.Config) {
	ingestor := climateMeasurements.NewIngestor(postgresConfig, kafkaConfig)
	defer ingestor.Close()

	for {
		ingestor.Execute()

		time.Sleep(5 * time.Second)
	}
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
