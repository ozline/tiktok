package cache

import (
	"context"
	"strconv"

	"github.com/ozline/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

func IsVideoLikeExist(ctx context.Context, videoID int64, userID int64) (bool, error) {
	exist, err := RedisClient.SIsMember(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Result()
	if err != nil {
		return exist, err
	}
	return exist, nil
}

func AddVideoLikeCount(ctx context.Context, videoID int64, userID int64) error {
	pipe := RedisClient.TxPipeline()

	// add user like
	if err := pipe.SAdd(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		return err
	}

	// add video like count
	if err := pipe.Incr(ctx, GetVideoLikeCountKey(videoID)).Err(); err != nil {
		return err
	}

	// expire video like count key
	if err := pipe.Expire(ctx, GetVideoLikeCountKey(videoID), constants.LikeExpiredTime).Err(); err != nil {
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
	// delete user like
	if err := pipe.SRem(ctx, GetUserLikeKey(userID), strconv.FormatInt(videoID, 10)).Err(); err != nil {
		return err
	}

	// reduce video like count
	if err := pipe.Decr(ctx, GetVideoLikeCountKey(videoID)).Err(); err != nil {
		return err
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func SetVideoLikeCount(ctx context.Context, videoID int64, count int64) error {
	err := RedisClient.Set(ctx, GetVideoLikeCountKey(videoID), count, constants.LikeExpiredTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetVideoLikeCount(ctx context.Context, videoID int64) (bool, int64, error) {
	count, err := RedisClient.Get(ctx, GetVideoLikeCountKey(videoID)).Result()
	if err == redis.Nil {
		return false, 0, nil
	}
	if err != nil {
		return true, 0, err
	}

	likeCount, _ := strconv.ParseInt(count, 10, 64)
	return true, likeCount, nil
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
