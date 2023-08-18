package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, error) {
	return db.GetVideoInfoByTime(s.ctx, req.LatestTime)
}
