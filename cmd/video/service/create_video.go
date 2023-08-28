package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func (s *VideoService) CreateVideo(req *video.PutVideoRequest, playURL string, coverURL string) (*db.Video, error) {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, errno.AuthorizationFailedError
	}
	videoModel := &db.Video{
		UserID:   claim.UserId,
		PlayUrl:  playURL,
		CoverUrl: coverURL,
		Title:    req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}
