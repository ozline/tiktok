package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, error) {
	if exist, err := cache.IsExistVideoInfo(s.ctx, req.LatestTime); exist == 1 {
		if err != nil {
			return nil, err
		} else {
			return cache.GetVideoList(s.ctx, req.LatestTime)
		}
	} else {
		formattedTime := time.Unix(req.LatestTime, 0).Format("2006-01-02 15:04:05")
		videoList, err := db.GetVideoInfoByTime(s.ctx, formattedTime)
		go cache.AddVideoList(s.ctx, videoList, req.LatestTime)
		return videoList, err
	}

}
