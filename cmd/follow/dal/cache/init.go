package cache

import (
	"context"

	"github.com/ozline/tiktok/config"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	//redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       0,
	})

	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}
