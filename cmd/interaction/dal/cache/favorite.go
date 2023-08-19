package cache

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

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
	VideoExist, err := RedisClient.SIsMember(ctx, GetVideoKey(videoId), strconv.FormatInt(userId, 10)).Result()
	if err != nil || !VideoExist {
		klog.Infof("video no exist")
		return err
	}
	if RedisClient.SRem(ctx, GetVideoKey(videoId), strconv.FormatInt(userId, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}

	userExist, err := RedisClient.SIsMember(ctx, GetUserKey(userId), strconv.FormatInt(videoId, 10)).Result()
	if err != nil || !userExist {
		klog.Infof("user no exist")
		return err
	}
	if RedisClient.SRem(ctx, GetUserKey(userId), strconv.FormatInt(videoId, 10)).Err(); err != nil {
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
	}

	videosId := make([]int64, 10)
	for _, item := range items {
		videoId, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			klog.Infof("parseInt err")
			return nil, err
		}
		videosId = append(videosId, videoId)
	}
	return videosId, nil
}
