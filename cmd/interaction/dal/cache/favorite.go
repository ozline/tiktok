package cache

import (
	"context"
	"strconv"
)

func IsVideoLikeExist(ctx context.Context, videoID int64, userID int64) (bool, error) {
	exist, err := RedisClient.SIsMember(ctx, GetVideoLikeCountKey(videoID), strconv.FormatInt(userID, 10)).Result()
	if err != nil {
		return exist, err
	}
	return exist, nil
}

func AddVideoLikeCount(ctx context.Context, videoID int64, userID int64) error {
	pipe := RedisClient.TxPipeline()
	// add video like
	if err := pipe.SAdd(ctx, GetVideoLikeCountKey(videoID), strconv.FormatInt(userID, 10)).Err(); err != nil {
		return err
	}
	// add user like
	if err := pipe.SAdd(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		return err
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func ReduceVideoLikeCount(ctx context.Context, videoID int64, userID int64) error {
	pipe := RedisClient.TxPipeline()
	// unlike the video
	if err := pipe.SRem(ctx, GetVideoLikeCountKey(videoID), strconv.FormatInt(userID, 10)).Err(); err != nil {
		return err
	}
	if err := pipe.SRem(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		return err
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func GetVideoLikeCount(ctx context.Context, videoID int64) (int64, error) {
	count, err := RedisClient.SCard(ctx, GetVideoLikeCountKey(videoID)).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserFavoriteVideos(ctx context.Context, userID int64) ([]int64, error) {
	items, err := RedisClient.SMembers(ctx, GetUserLikeKey(userID)).Result()
	if err != nil {
		return nil, err
	}

	// get favorite video id list
	videoIDList := make([]int64, 0, len(items))
	for _, item := range items {
		videoId, _ := strconv.ParseInt(item, 10, 64)
		videoIDList = append(videoIDList, videoId)
	}
	return videoIDList, nil
}

func UpdateFavoriteVideoList(ctx context.Context, userID int64, videoIDList []int64) error {
	var err error
	for _, videoID := range videoIDList {
		err = RedisClient.SAdd(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Err()
	}
	return err
}

func GetUserFavoriteCount(ctx context.Context, userID int64) (int64, error) {
	count, err := RedisClient.SCard(ctx, GetUserLikeKey(userID)).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
