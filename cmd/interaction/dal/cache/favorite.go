package cache

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

func IsVideoLikeExist(ctx context.Context, videoId int64, userId int64) (bool, error) {
	exist, err := RedisClient.SIsMember(ctx, GetVideoKey(videoId), strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return true, nil
}

func AddVideoLikeCount(ctx context.Context, videoId int64, userId int64) error {
	// add video like
	if err := RedisClient.SAdd(ctx, GetVideoKey(videoId), strconv.FormatInt(userId, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	// add user like
	if err := RedisClient.SAdd(ctx, GetUserKey(userId), strconv.FormatInt(videoId, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func ReduceVideoLikeCount(ctx context.Context, videoId int64, userId int64) error {
	// unlike the video
	if err := RedisClient.SRem(ctx, GetVideoKey(videoId), strconv.FormatInt(userId, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	if err := RedisClient.SRem(ctx, GetUserKey(userId), strconv.FormatInt(videoId, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
	}
	return nil
}

func GetVideoLikeCount(ctx context.Context, videoId int64) (int64, error) {
	count, err := RedisClient.SCard(ctx, GetVideoKey(videoId)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return 0, err
	}
	return count, nil
}

func GetUserFavoriteVideos(ctx context.Context, userId int64) ([]int64, error) {
	items, err := RedisClient.SMembers(ctx, GetUserKey(userId)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	//get favorite video id list
	videoIdList := make([]int64, 10)
	for _, item := range items {
		videoId, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			klog.Infof("parseInt err")
			return nil, err
		}
		videoIdList = append(videoIdList, videoId)
	}
	return videoIdList, nil
}

func UpdateFavoriteVideoList(ctx context.Context, userId int64, videoIdList []int64) error {
	err := RedisClient.SAdd(ctx, GetUserKey(userId), videoIdList).Err()
	if err != nil {
		return err
	}
	return nil
}
