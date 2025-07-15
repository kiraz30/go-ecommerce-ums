package helpers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_HOST", "localhost:6379"),
		DB:   0,
	})

	//test connect redis
	ping, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		Logger.Error("Failed to connect Redis:", err)
		return
	}
	Logger.Info("PING REDIS :", ping)

	RedisClient = rdb
}
