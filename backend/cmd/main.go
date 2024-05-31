package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/rs/homecontrol/pkg/config"
	"github.com/rs/homecontrol/pkg/httpfetcher"
	"github.com/rs/homecontrol/pkg/repository"
	"github.com/rs/homecontrol/pkg/rest"
)

func main() {
	log.Println("Starting server")

	cfg := bootstrapConfig()

	bootstrapRepository(cfg)

	bootstrapHttpFetcher(cfg)

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
	repository.CreateInstance(cfg)
}

func bootstrapHttpFetcher(cfg *config.Config) {
	httpfetcher.RunService(cfg.HouseServer.Url, cfg.DataPaths.LatestMeasurements)
}
