package cache

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/video/dal/db"
)

func AddVideoList(ctx context.Context, videoList []db.Video, latestTime int64) {
	videoJson, err := json.Marshal(videoList)
	if err != nil {
		klog.Error(err)
	}
	err = RedisClient.Set(ctx, strconv.FormatInt(latestTime, 10), videoJson, time.Minute*10).Err()
	if err != nil {
		klog.Error(err)
	}
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
	exist, err = RedisClient.Exists(ctx, strconv.FormatInt(latestTime, 10)).Result()
	return
}
