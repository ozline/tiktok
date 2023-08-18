package cache

import (
	"errors"

	"github.com/ozline/tiktok/pkg/constants"
	redis "github.com/redis/go-redis/v9"
)

var (
	RedisDB *redis.Client
)

func Init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPWD,     // no password set
		DB:       constants.ReidsDB_Chat, // use default DB
	})
	//docker run -d --privileged=true -p 6379:6379 -v /usr/local/redis/conf/redis.conf:/etc/redis/redis.conf -v /usr/local/redis/data:/data --name redis-1 redis:latest redis-server /etc/redis/redis.conf --appendonly yes
	if RedisDB == nil {
		panic(errors.New("[redis init error]"))
	}
}
