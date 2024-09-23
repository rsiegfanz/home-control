package main

import (
	"fmt"
	"log"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/configs"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg"
)

func main() {
	config, err := pkg.LoadConfig[configs.FetcherConfig]()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Printf("Fetcher-Config: %+v\n", config)
}
