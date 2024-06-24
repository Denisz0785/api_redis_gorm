package database

import (
	"fmt"
	"redis_gorm_fiber/config"

	"github.com/go-redis/redis/v8"
)

func ConnectionRedisDB(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	fmt.Println("Connection to Redis DB success")

	return rdb
}
