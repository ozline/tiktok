package cache

import (
	"context"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

func MessageInsert(ctx context.Context, key string, revkey string, score float64, member string) error {
	// 先判断是否key是否存在，如果存在则判断过期时间是否小于十天，小于则加时间，大于则不加时间
	if ok := MessageExist(ctx, key); ok != 0 {
		err := RedisDB.ZAdd(context.TODO(), key, redis.Z{
			Score:  score,
			Member: member,
		}).Err()
		if err != nil {
			return err
		}
		lastTime, err := RedisDB.TTL(ctx, key).Result()
		if err != nil {
			return err
		}
		if lastTime.Seconds() < 60*60*24*10 {
			err = RedisDB.Expire(ctx, key, time.Hour*24*30).Err()
			if err != nil {
				return err
			}
			err = RedisDB.Expire(ctx, revkey, time.Hour*24*30).Err()
			if err != nil {
				return err
			}
			return nil
		}
	}
	// 不存在则直接add然后加时间
	err := RedisDB.ZAdd(context.TODO(), key, redis.Z{
		Score:  score,
		Member: member,
	}).Err()
	if err != nil {
		return err
	}
	err = RedisDB.Expire(ctx, key, time.Hour*24*30).Err()
	if err != nil {
		return err
	}
	err = RedisDB.Expire(ctx, revkey, time.Hour*24*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func MessageGet(ctx context.Context, key string) ([]string, error) {
	mem, err := RedisDB.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
		Min: strconv.Itoa(0),
		Max: strconv.Itoa(int(time.Now().Unix())),
	}).Result()
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func MessageExist(ctx context.Context, key string) int64 {
	ok, err := RedisDB.Exists(ctx, key).Result()
	if err != nil {
		return 0
	}
	return ok
}
