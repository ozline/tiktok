package cache

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
	"strconv"
)

func GetComments(ctx context.Context, videoId int64) (comments *[]redis.Z, err error) {
	rComments, err := RedisClient.ZRevRangeWithScores(ctx, strconv.FormatInt(videoId, 10), 0, -1).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return nil, err
	}
	klog.Infof("Get comments : videoId %v: %v\n", videoId, rComments)
	return &rComments, nil
}

func AddComment(ctx context.Context, videoId int64, comment *db.Comment) (err error) {
	//json序列化
	data, err := (*comment).MarshalMsg(nil)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return
	}
	err = RedisClient.ZAdd(ctx, strconv.FormatInt(videoId, 10), redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data}).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comment: videoId %v comment %v date %v\n", videoId, comment)
	}
	return
}

func AddComments(ctx context.Context, videoId int64, comments *[]db.Comment) (err error) {
	var zComments []redis.Z
	for _, comment := range *comments {
		data, err := comment.MarshalMsg(nil)
		if err != nil {
			klog.Infof("Error: %v\n", err)
			return err
		}
		zComments = append(zComments, redis.Z{Score: float64(comment.CreatedAt.Unix()), Member: data})
	}
	a, err := RedisClient.ZAdd(ctx, strconv.FormatInt(videoId, 10), zComments...).Result()
	fmt.Println(a)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return err
	}
	err = RedisClient.Expire(ctx, strconv.FormatInt(videoId, 10), constants.CommentExpiredTime).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comments: videoId %v \n", strconv.FormatInt(videoId, 10))
	}
	return
}

func DeleteComment(ctx context.Context, videoId int64, comment *db.Comment) (err error) {
	data, err := (*comment).MarshalMsg(nil)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return
	}
	err = RedisClient.ZRem(ctx, strconv.FormatInt(videoId, 10), data).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Delete comment: videoId %v comment %v \n", videoId, comment)
	}
	return
}

func CountComments(ctx context.Context, videoId int64) (count int64, err error) {
	count, err = RedisClient.ZCard(ctx, strconv.FormatInt(videoId, 10)).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Count comment: videoId %v count %v \n", videoId, count)
	}
	return
}

func IsExistComment(ctx context.Context, videoId int64) (exist int64, err error) {
	fmt.Println()
	exist, err = RedisClient.Exists(ctx, strconv.FormatInt(videoId, 10)).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Is exist comment: videoId %v exist %v \n", videoId, exist)
	}
	return
}
