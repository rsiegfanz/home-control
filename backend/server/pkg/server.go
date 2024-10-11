package pkg

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rsiegfanz/home-control/backend/server/pkg/configs"
	graphqlHandlers "github.com/rsiegfanz/home-control/backend/server/pkg/graphql/handlers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/resolvers"
	restHandlers "github.com/rsiegfanz/home-control/backend/server/pkg/rest/handlers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/middleware"
	"gorm.io/gorm"
)

func NewServer(serverConfig configs.ServerConfig, db *gorm.DB) *http.Server {
	router := mux.NewRouter()

	registerRest(router, db)
	registerGraphQl(router, db)
	registerStaticContent(router)

	router.Use(middleware.LoggingMiddleware)

	cors := cors.AllowAll()
	handler := cors.Handler(router)

	srv := &http.Server{
		Addr:         serverConfig.Adress,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handler,
	}

	return srv
}

func registerRest(router *mux.Router, db *gorm.DB) {
	restHandlers := restHandlers.NewHandler(db)

	// router.HandleFunc("/", controller.HelloWorldHandler)

	router.HandleFunc("/rooms", restHandlers.GetRoomsHandler)

	router.HandleFunc("/rooms/{roomId}/measurements", restHandlers.GetMeasurementsByRoomIdHandler)
	router.HandleFunc("/rooms/{roomId}/measurements/latest", restHandlers.GetLatestMeasurementByRoomIdHandler)
}

func registerGraphQl(router *mux.Router, db *gorm.DB) {
	queryResolver := &resolvers.QueryResolver{DB: db}
	router.HandleFunc("/graphql", graphqlHandlers.NewGraphQLHandler(queryResolver)).Methods("POST")
}

func registerStaticContent(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/")))
}
