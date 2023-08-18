package cache

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/config"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init() {
	ctx := context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       2,
	})
	_, err := Rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	err = Rdb.Set(ctx, "test", "just for test", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := Rdb.Get(ctx, "test").Result()

	if err != nil {
		panic(err)
	}

	klog.Infof("val: %v\n", val)
}
