package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, error) {
	formattedTime := time.Unix(req.LatestTime, 0).Format("2006-01-02 15:04:05")
	return db.GetVideoInfoByTime(s.ctx, formattedTime)
}
