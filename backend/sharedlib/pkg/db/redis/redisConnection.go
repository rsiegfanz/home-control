package redis

import (
	"github.com/redis/go-redis/v9"
)

var groupId = "rs"

func InitClient(config Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: config.Host,
	})
}
