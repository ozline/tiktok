package test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	redis "github.com/redis/go-redis/v9"
)

func testRedis(t *testing.T) {
	cache.Init()
	res, err := cache.RedisDB.Exists(context.Background(), "key").Result()
	if res != 0 {
		mem, err := cache.RedisDB.ZRevRangeByScore(context.Background(), "key", &redis.ZRangeBy{
			Min: "0",
			Max: strconv.FormatInt(time.Now().Unix(), 10),
		}).Result()
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		t.Logf("result--->%v", mem)
	} else if err != nil {
		t.Error(err)
		t.Fail()
	}
}
