package cache

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
)

func IsVideoLikeExist(ctx context.Context, videoID int64, userID int64) (bool, error) {
	exist, err := RedisClient.SIsMember(ctx, GetVideoKey(videoID), strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return exist, err
	}
	return exist, nil
}

func AddVideoLikeCount(ctx context.Context, videoID int64, userID int64) error {
	// add video like
	if err := RedisClient.SAdd(ctx, GetVideoKey(videoID), strconv.FormatInt(userID, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	// add user like
	if err := RedisClient.SAdd(ctx, GetUserKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func ReduceVideoLikeCount(ctx context.Context, videoID int64, userID int64) error {
	// unlike the video
	if err := RedisClient.SRem(ctx, GetVideoKey(videoID), strconv.FormatInt(userID, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	if err := RedisClient.SRem(ctx, GetUserKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		klog.Infof("err: %v", err)
	}
	return nil
}

func GetVideoLikeCount(ctx context.Context, videoID int64) (int64, error) {
	count, err := RedisClient.SCard(ctx, GetVideoKey(videoID)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return 0, err
	}
	return count, nil
}

func GetUserFavoriteVideos(ctx context.Context, userID int64) ([]int64, error) {
	items, err := RedisClient.SMembers(ctx, GetUserKey(userID)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return nil, err
	}

	//get favorite video id list
	videoIDList := make([]int64, 0, 10)
	for _, item := range items {
		videoId, _ := strconv.ParseInt(item, 10, 64)
		videoIDList = append(videoIDList, videoId)
	}
	return videoIDList, nil
}

func UpdateFavoriteVideoList(ctx context.Context, userID int64, videoIDList []int64) error {
	return RedisClient.SAdd(ctx, GetUserKey(userID), videoIDList).Err()
}
