package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/ozline/tiktok/cmd/video/dal/db"
)

func AddVideoList(ctx context.Context, videoList []db.Video, latestTime int64) (err error) {
	videoJson, err := json.Marshal(videoList)
	if err != nil {
		return
	}
	err = RedisClient.Set(ctx, strconv.FormatInt(latestTime, 10), videoJson, 0).Err()
	return
}
func GetVideoList(ctx context.Context, latestTime int64) (videoList []db.Video, err error) {
	data, err := RedisClient.Get(ctx, strconv.FormatInt(latestTime, 10)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &videoList)
	if err != nil {
		return nil, err
	}
	return
}
func IsExistVideoInfo(ctx context.Context, latestTime int64) (exist int64, err error) {
	fmt.Println(strconv.FormatInt(latestTime, 10))
	exist, err = RedisClient.Exists(ctx, strconv.FormatInt(latestTime, 10)).Result()
	return
}
