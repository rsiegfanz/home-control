module github.com/rsiegfanz/home-control/backend/server

go 1.23.1

require (
	github.com/rsiegfanz/home-control/backend/sharedlib v0.0.0
	github.com/gorilla/mux v1.8.1
	go.uber.org/zap v1.27.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/gorilla/handlers v1.5.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/rs/cors v1.10.1 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/rsiegfanz/home-control/backend/sharedlib => ../sharedlib
