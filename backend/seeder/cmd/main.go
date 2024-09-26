package main

import (
	"log"

	"github.com/rsiegfanz/home-control/backend/fetcher/pkg/seeds"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/config"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres"
	"github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/postgres/models"
)

func main() {
	log.Printf("Seeder started...")

	config := config.ConfigPostgres{}
	config.Host = "localhost"
	config.Port = 5432
	config.DbName = "home_control_db"
	config.User = "home_control_user"
	config.Password = "home_control_password"

	/*config, err := config.LoadConfig[config.ConfigPostgres]()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}*/

	log.Printf("Seeder-Config: %+v\n", config)

	db, err := postgres.InitDB(config)

	rooms := []models.Room{}
	db.Find(&rooms)

	log.Printf("rooms: ", rooms)

	err = seeds.SeedRooms(db)
	if err != nil {
		log.Fatalf("Seed error %v", config)
	}

	log.Printf("Finished seeds")
}
