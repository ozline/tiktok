package cache

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

func GetComments(ctx context.Context, key string) (comments *[]redis.Z, err error) {
	pipe := RedisClient.TxPipeline()
	err = pipe.TTL(ctx, key).Err()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	err = pipe.ZRevRangeWithScores(ctx, key, 0, -1).Err()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	for _, cmder := range cmders {
		err = cmder.Err()
		if err != nil {
			klog.Error(err)
			return nil, err
		}
	}
	lastTime := cmders[0].(*redis.DurationCmd).Val()
	rComments := cmders[1].(*redis.ZSliceCmd).Val()
	if lastTime < constants.CommentExpiredTime/2 {
		err = RedisClient.Expire(ctx, key, constants.CommentExpiredTime).Err()
		if err != nil {
			klog.Error(err)
			return nil, err
		}
	}
	klog.Infof("Get comments : videoId %v\n", key)
	return &rComments, nil
}

func AddComment(ctx context.Context, key string, comment *db.Comment) (err error) {
	data, err := comment.MarshalMsg(nil)
	if err != nil {
		klog.Error(err)
		return
	}
	err = RedisClient.ZAdd(ctx, key, redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data}).Err()
	if err != nil {
		klog.Error(err)
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
			klog.Error(err)
			return err
		}
		zComments = append(zComments, redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data})
	}
	pipe := RedisClient.TxPipeline()
	err = pipe.ZAdd(ctx, key, zComments...).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	err = pipe.Expire(ctx, key, constants.CommentExpiredTime).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		klog.Error(err)
		return err
	}
	for _, cmder := range cmders {
		err = cmder.Err()
		if err != nil {
			klog.Error(err)
			return err
		}
	}
	klog.Infof("Add comments: videoId %v \n", key)
	return err
}

func DeleteComment(ctx context.Context, key string, comment *db.Comment) (err error) {
	data, err := comment.MarshalMsg(nil)
	if err != nil {
		klog.Error(err)
		return
	}
	err = RedisClient.ZRem(ctx, key, data).Err()
	if err != nil {
		klog.Error(err)
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
		klog.Error(err)
	} else {
		klog.Infof("Count comment: videoId %v count %v \n", key, count)
	}
	return
}

func IsExistComment(ctx context.Context, key string) (exist int64, err error) {
	exist, err = RedisClient.Exists(ctx, key).Result()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Is exist comment: videoId %v exist %v \n", key, exist)
	}
	return
}

func DeleteComments(ctx context.Context, key string) (err error) {
	err = RedisClient.Del(ctx, key).Err()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Delete comments: videoId %v \n", key)
	}
	return
}
