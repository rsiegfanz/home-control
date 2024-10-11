package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	. "github.com/rsiegfanz/home-control/backend/server/pkg"
	"github.com/rsiegfanz/home-control/backend/server/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	logPath := "/var/log"
	if !config.IsProd() {
		logPath = "d:\\dev\\docker\\share\\home-control\\promtail"
	}

	if err := logging.InitLogger("info", "server", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Server started")

	dbConfig, serverConfig := loadConfigs()
	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	server := NewServer(serverConfig, db)

	go func() {
		log.Println("Start")
		if err := server.ListenAndServe(); err != nil {
			logging.Logger.Error("Error starting server", zap.Error(err))
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	logging.Logger.Debug("Stopping server")

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	server.Shutdown(ctx)

	logging.Logger.Info("Server stopped")
}

func loadConfigs() (postgres.Config, configs.ServerConfig) {
	if config.IsProd() {
		return loadConfigsProd()
	}
	return loadConfigsDev()
}

func loadConfigsProd() (postgres.Config, configs.ServerConfig) {
	logging.Logger.Info("Loading PROD environment")

	postgresConfig, err := config.LoadConfig[postgres.Config]()
	if err != nil {
		logging.Logger.Fatal("Error loading postgres config", zap.Error(err))
	}

	serverConfig, err := config.LoadConfig[configs.ServerConfig]()
	if err != nil {
		logging.Logger.Fatal("Error loading config", zap.Error(err))
	}

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, serverConfig
}

func loadConfigsDev() (postgres.Config, configs.ServerConfig) {
	logging.Logger.Warn("Loading DEV environment")

	postgresConfig := postgres.Config{}
	postgresConfig.Host = "localhost"
	postgresConfig.Port = 5432
	postgresConfig.DbName = "home_control_db"
	postgresConfig.User = "home_control_user"
	postgresConfig.Password = "home_control_password"

	serverConfig := configs.ServerConfig{}
	serverConfig.Adress = "0.0.0.0:8080"

	logging.Logger.Debug("Configs loaded")

	return postgresConfig, serverConfig
}
