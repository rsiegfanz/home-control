package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/controllers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/middleware"
	"gorm.io/gorm"
)

func NewServer(db *gorm.DB) *http.Server {
	router := mux.NewRouter()

	controller := controllers.NewController(db)

	// router.HandleFunc("/", controller.HelloWorldHandler)

	router.HandleFunc("/rooms", controller.GetRoomsHandler)

	router.HandleFunc("/rooms/{roomId}/measurements", controller.GetMeasurementsByRoomIdHandler)
	router.HandleFunc("/rooms/{roomId}/measurements/latest", controller.GetLatestMeasurementByRoomIdHandler)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/")))

	router.Use(middleware.LoggingMiddleware)

	cors := cors.AllowAll()
	handler := cors.Handler(router)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	return srv
}
