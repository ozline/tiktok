package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type Comment struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
}

func GetComments(ctx context.Context, videoId int64) (comments *[]redis.Z, err error) {
	rComments, err := RedisClient.ZRevRangeWithScores(ctx, strconv.FormatInt(videoId, 10), 0, -1).Result()
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return nil, err
	}
	klog.Infof("Get comments : videoId %v: %v\n", videoId, rComments)
	return &rComments, nil
}

func AddComment(ctx context.Context, videoId int64, comment *Comment, date float64) (err error) {
	//json序列化
	data, err := json.Marshal(*comment)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return
	}
	err = RedisClient.ZAdd(ctx, strconv.FormatInt(videoId, 10), redis.Z{Score: date, Member: data}).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comment: videoId %v comment %v date %v\n", videoId, data, date)
	}
	return
}

func AddComments(ctx context.Context, videoId int64, comments *[]Comment, date *[]float64) (err error) {
	var zComments []redis.Z
	for index, comment := range *comments {
		data, err := json.Marshal(comment)
		if err != nil {
			klog.Infof("Error: %v\n", err)
			return err
		}
		zComments = append(zComments, redis.Z{Score: (*date)[index], Member: data})
	}
	a, err := RedisClient.ZAdd(ctx, strconv.FormatInt(videoId, 10), zComments...).Result()
	fmt.Println(a)
	if err != nil {
		klog.Infof("Error: %v\n", err)
		return err
	}
	err = RedisClient.Expire(ctx, strconv.FormatInt(videoId, 10), 60*time.Minute).Err()
	if err != nil {
		klog.Infof("Error: %v\n", err)
	} else {
		klog.Infof("Add comments: videoId %v \n", strconv.FormatInt(videoId, 10))
	}
	return
}

func DeleteComment(ctx context.Context, videoId int64, comment *Comment) (err error) {
	data, err := json.Marshal(*comment)
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
