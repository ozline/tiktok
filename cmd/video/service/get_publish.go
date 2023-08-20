package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) GetPublishVideoInfo(req *video.GetPublishListRequest) ([]db.Video, error) {
	return db.GetVideoInfoByUid(s.ctx, req.UserId)
}
