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
	t.Log("----------BEGIN------------")
	cache.Init()
	res, err := cache.RedisDB.Exists(context.Background(), "key1").Result()
	if res != 0 {
		mem, err := cache.RedisDB.ZRevRangeByScore(context.Background(), "key1", &redis.ZRangeBy{
			Min: strconv.Itoa(0),
			Max: strconv.Itoa(int(time.Now().Unix())),
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
	t.Log("----------SUCCESS----------")
}
