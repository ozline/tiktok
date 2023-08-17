package cache

import (
	"context"

	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	//redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr: constants.RedisAddr,
		// Password: constants.RedisPWD,
		DB: 2, //constants.RedisDBFollow
	})
	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
}
