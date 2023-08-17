package cache

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

// A sample for redis

func Init() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddr,
		Password: constants.RedisPassword,
		DB:       0,
	})

	err := rdb.Set(ctx, "test", "just for test", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "test").Result()

	if err != nil {
		panic(err)
	}

	klog.Infof("val: %v\n", val)

	val, err = rdb.Get(ctx, "test1").Result()

	if err == redis.Nil {
		klog.Info("Not found test1 key")
	} else if err != nil {
		panic(err)
	} else {
		klog.Infof("val: %v\n", val)
	}
}
