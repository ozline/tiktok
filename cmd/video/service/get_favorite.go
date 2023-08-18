package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) GetFavoriteVideoInfo(req *video.GetFavoriteVideoInfoRequest) ([]db.Video, error) {
	return db.GetVideoInfoByID(s.ctx, req.VideoId)
}
