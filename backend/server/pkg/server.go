package pkg

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"github.com/rsiegfanz/home-control/backend/server/pkg/configs"
	graphqlHandlers "github.com/rsiegfanz/home-control/backend/server/pkg/graphql/handlers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/graphql/resolvers"
	restHandlers "github.com/rsiegfanz/home-control/backend/server/pkg/rest/handlers"
	"github.com/rsiegfanz/home-control/backend/server/pkg/rest/middleware"
	rsredis "github.com/rsiegfanz/home-control/backend/sharedlib/pkg/db/redis"
	"gorm.io/gorm"
)

func NewServer(serverConfig configs.ServerConfig, db *gorm.DB, redisConfig rsredis.Config) *http.Server {
	router := mux.NewRouter()

	redisClient := rsredis.InitClient(redisConfig)

	registerRest(router, db)
	registerGraphQl(router, db, redisClient)
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

func registerGraphQl(router *mux.Router, db *gorm.DB, redisClient *redis.Client) {
	queryResolver := &resolvers.QueryResolver{DB: db, RedisClient: redisClient}
	router.HandleFunc("/graphql", graphqlHandlers.NewGraphQLHandler(queryResolver)).Methods("GET", "POST")
}

func registerStaticContent(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp/")))
}
