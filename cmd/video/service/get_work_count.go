package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func (s *VideoService) GetWorkCount(req *video.GetWorkCountRequest) (workCount int64, err error) {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return 0, errno.AuthorizationFailedError
	}
	return db.GetWorkCountByUid(s.ctx, claim.UserId)
}
