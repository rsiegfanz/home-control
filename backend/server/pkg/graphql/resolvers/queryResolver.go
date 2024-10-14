package resolvers

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type QueryResolver struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}
