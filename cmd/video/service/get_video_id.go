package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func (s *VideoService) GetVideoIDByUid(req *video.GetVideoIDByUidRequset) (videoIDList []int64, err error) {
	return db.GetVideoIDByUid(s.ctx, req.UserId)
}
