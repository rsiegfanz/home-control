package presentation

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/homecontrol/pkg/presentation/controllers"
	"github.com/rs/homecontrol/pkg/presentation/middleware"
)

func NewServer() *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HelloWorldHandler)

	router.HandleFunc("/rooms", controllers.GetRoomsHandler)

	router.HandleFunc("/rooms/{roomId}/temperatures", controllers.GetTemperaturesByRoomIdHandler)
	router.HandleFunc("/rooms/{roomId}/temperatures/latest", controllers.GetLatestTemperatureByRoomIdHandler)

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
