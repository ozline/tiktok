package cache

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

func GetComments(ctx context.Context, key string) (comments *[]redis.Z, err error) {
	lastTime, err := RedisClient.TTL(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	if lastTime < constants.CommentExpiredTime/2 {
		err = RedisClient.Expire(ctx, key, constants.CommentExpiredTime).Err()
		if err != nil {
			return nil, err
		}
	}
	rComments, err := RedisClient.ZRevRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return nil, err
	}
	klog.Infof("Get comments : videoId %v: %v\n", key, rComments)
	return &rComments, nil
}

func AddComment(ctx context.Context, key string, comment *db.Comment) (err error) {
	data, err := (*comment).MarshalMsg(nil)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return
	}
	err = RedisClient.ZAdd(ctx, key, redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data}).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comment: videoId %v comment %v date %v\n", key, comment)
	}
	return
}

func AddComments(ctx context.Context, key string, comments *[]db.Comment) (err error) {
	var zComments []redis.Z
	for _, comment := range *comments {
		data, err := comment.MarshalMsg(nil)
		if err != nil {
			klog.Infof("Error: %v\n", err)
			return err
		}
		zComments = append(zComments, redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data})
	}
	err = RedisClient.ZAdd(ctx, key, zComments...).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return err
	}
	err = RedisClient.Expire(ctx, key, constants.CommentExpiredTime).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comments: videoId %v \n", key)
	}
	return
}

func DeleteComment(ctx context.Context, key string, comment *db.Comment) (err error) {
	data, err := (*comment).MarshalMsg(nil)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return
	}
	err = RedisClient.ZRem(ctx, key, data).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Delete comment: videoId %v comment %v \n", key, comment)
	}
	return
}

func CountComments(ctx context.Context, key string) (count int64, err error) {
	lastTime, err := RedisClient.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	if lastTime < constants.CommentExpiredTime/2 {
		err = RedisClient.Expire(ctx, key, constants.CommentExpiredTime).Err()
		if err != nil {
			return 0, err
		}
	}
	count, err = RedisClient.ZCard(ctx, key).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Count comment: videoId %v count %v \n", key, count)
	}
	return
}

func IsExistComment(ctx context.Context, key string) (exist int64, err error) {
	exist, err = RedisClient.Exists(ctx, key).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Is exist comment: videoId %v exist %v \n", key, exist)
	}
	return
}
