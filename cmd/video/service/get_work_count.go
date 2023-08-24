package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func (s *VideoService) GetWorkCount(req *video.GetWorkCountRequest) (workCount int64, err error) {
	return db.GetWorkCountByUid(s.ctx, req.UserId)
}
