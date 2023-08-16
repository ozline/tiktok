package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) GetVideoInfo(req *video.GetFavoriteVideoInfoRequest) ([]db.Video, error) {
	return db.GetVideoInfo(s.ctx, req.VideoId)
}
