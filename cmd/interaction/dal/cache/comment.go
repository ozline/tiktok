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
	commentKey := GetCommentKey(key)
	err = pipe.TTL(ctx, commentKey).Err()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	err = pipe.ZRevRangeWithScores(ctx, commentKey, 0, -1).Err()
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
		err = RedisClient.Expire(ctx, commentKey, constants.CommentExpiredTime).Err()
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
	commentKey := GetCommentKey(key)
	zComments := make([]redis.Z, len(*comments))
	for i := 0; i < len(*comments); i++ {
		data, err := (*comments)[i].MarshalMsg(nil)
		if err != nil {
			klog.Error(err)
			return err
		}
		zComments[i] = redis.Z{Score: float64((*comments)[i].CreatedAt.Unix()), Member: data}
	}
	pipe := RedisClient.TxPipeline()
	err = pipe.ZAdd(ctx, commentKey, zComments...).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	err = pipe.Expire(ctx, commentKey, constants.CommentExpiredTime).Err()
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

func AddNoData(ctx context.Context, key string) (err error) {
	zData := redis.Z{}
	pipe := RedisClient.TxPipeline()
	commentKey := GetCommentKey(key)
	err = pipe.ZAdd(ctx, commentKey, zData).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	err = pipe.Expire(ctx, commentKey, constants.NoDataExpiredTime).Err()
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
	klog.Infof("Add NoData: videoId %v \n", key)
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

func GetCount(ctx context.Context, key string) (ok bool, count string, err error) {
	count, err = RedisClient.Get(ctx, GetCountKey(key)).Result()
	if err == redis.Nil {
		klog.Infof("Count comment: videoId %v count nil ! \n", key)
		return false, count, nil
	}
	if err == nil {
		klog.Infof("Count comment: videoId %v count %v \n", key, count)
	} else {
		klog.Error(err)
	}
	return true, count, err
}

func SetCount(ctx context.Context, key string, count int64) (err error) {
	err = RedisClient.Set(ctx, GetCountKey(key), count, constants.CommentExpiredTime).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	klog.Infof("Set Count: videoId %v \n", key)
	return err
}

func IsExistComment(ctx context.Context, key string) (exist int64, err error) {
	commentKey := GetCommentKey(key)
	exist, err = RedisClient.Exists(ctx, commentKey).Result()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Is exist comment: videoId %v exist %v \n", key, exist)
	}
	return
}

func Delete(ctx context.Context, key string) (err error) {
	err = RedisClient.Del(ctx, key).Err()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Delete : %v \n", key)
	}
	return
}

func Unlink(ctx context.Context, key string) (err error) {
	commentKey := GetCommentKey(key)
	err = RedisClient.Unlink(ctx, commentKey).Err()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Unlink : %v \n", key)
	}
	return
}

func Lock(ctx context.Context, key string) (ok bool, err error) {
	ok, err = RedisClient.SetNX(ctx, key, 1, constants.LockTime).Result()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Lock: set %v %v \n", key, ok)
	}
	return
}

func AddCount(ctx context.Context, increment int64, videoID string) (err error) {
	err = RedisClient.IncrBy(ctx, GetCountKey(videoID), increment).Err()
	if err != nil {
		klog.Error(err)
	} else {
		klog.Infof("Add count: videoId %v increment %v\n", videoID, increment)
	}
	return
}
