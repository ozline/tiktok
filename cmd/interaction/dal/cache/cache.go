package cache

import (
	"context"
	"github.com/ozline/tiktok/pkg/constants"
)

func AddVideoLikeCount(ctx context.Context, videoId string) error {
	if err := RedisClient.HIncrBy(ctx, constants.RedisHashName, GetVideoKey(videoId), 1).Err(); err != nil {
		return err
	}
	return nil
}

func ReduceVideoLikeCount(ctx context.Context, videoId string) error {
	if err := RedisClient.HIncrBy(ctx, constants.RedisHashName, GetVideoKey(videoId), -1).Err(); err != nil {
		return err
	}
	return nil
}

func GetLikeCount(ctx context.Context, videoId string) (int64, error) {
	likeCount, err := RedisClient.HGet(ctx, constants.RedisHashName, GetVideoKey(videoId)).Int64()
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}
