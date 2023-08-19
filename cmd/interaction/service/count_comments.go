package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) CountComments(req *interaction.CommentCountRequest) (count int64, err error) {
	videoId := req.VideoId

	exist, err := cache.IsExistComment(s.ctx, videoId)
	if err != nil {
		return 0, err
	}

	if exist == 1 {
		count, err = cache.CountComments(s.ctx, videoId)
	} else {
		count, err = db.CountCommentsByVideoID(s.ctx, videoId)
	}

	if err != nil {
		return 0, err
	}
	return
}
