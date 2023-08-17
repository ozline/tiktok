package test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	redis "github.com/redis/go-redis/v9"
)

func TestRedis(t *testing.T) {
	dal.Init()
	res, _ := cache.RedisDB.Exists(context.Background(), "key1").Result()
	fmt.Println(res)
	mem, err := cache.RedisDB.ZRevRangeByScore(context.Background(), "key1", &redis.ZRangeBy{
		Min: strconv.Itoa(0),
		Max: strconv.Itoa(int(time.Now().Unix())),
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(mem)
}
