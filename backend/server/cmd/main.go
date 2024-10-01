package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/logging"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
)

func main() {
	logPath := "d:\\dev\\docker\\share\\home-control\\server"
	// logPath := "/mnt/d/dev/docker/share"
	if err := logging.InitLogger("info", "server", logPath); err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	defer logging.SyncLogger()

	if logging.Logger == nil {
		log.Fatalf("Logger not initialized")
	}
	defer logging.SyncLogger()

	logging.Logger.Info("Server started")

	dbConfig := loadConfigs()
	db, err := postgres.InitDB(dbConfig)
	if err != nil {
		logging.Logger.Fatal("Error opening database", zap.Error(err))
	}

	logging.Logger.Info("Server stopped")
}

func loadConfigs() postgres.Config {
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

	logging.Logger.Debug("Configs loaded")

	return postgresConfig
}

/*
func main() {
	log.Println("Starting server")

	cfg := bootstrapConfig()

	bootstrapRepository(cfg)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := rest.NewServer()

	go func() {
		log.Println("Start")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	log.Println("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Stopped")
	os.Exit(0)
}

func bootstrapConfig() *config.Config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	log.Println("cwd: ", cwd)

	cfg, err := config.Read(path.Join(cwd, "config.yaml"))
	if err != nil {
		log.Fatal(err)
		os.Exit(-2)
	}

	log.Println("config: ", cfg)
	return cfg
}

func bootstrapRepository(cfg *config.Config) {
	_, err := repository.CreateInstance(cfg)
	if err != nil {
		log.Fatal("could not init repository: ", err)
		os.Exit(-3)
	}
}
*/
