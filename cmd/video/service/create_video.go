package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func (s *VideoService) CreateVideo(req *video.PutVideoRequest, playURL string, coverURL string) (*db.Video, error) {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, errno.AuthorizationFailedError
	}
	playUrl, ok := s.ctx.Value("playUrl").(string)
	if !ok {
		return nil, errno.ServiceInternalError
	}
	coverUrl, ok := s.ctx.Value("coverUrl").(string)
	if !ok {
		return nil, errno.ServiceInternalError
	}
	videoModel := &db.Video{
		UserID:   claim.UserId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}
