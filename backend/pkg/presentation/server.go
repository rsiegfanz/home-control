package presentation

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	controller "github.com/rs/homecontrol/pkg/presentation/controllers"
	"github.com/rs/homecontrol/pkg/presentation/middleware"
)

func NewServer() *http.Server {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middleware.LoggingMiddleware)

	router.HandleFunc("/", controller.HelloWorldHandler)
	router.HandleFunc("/rooms", controller.GetRoomsHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	return srv
}
