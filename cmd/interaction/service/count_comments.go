package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"strconv"
)

func (s *InteractionService) CountComments(req *interaction.CommentCountRequest) (count int64, err error) {
	videoId := req.VideoId

	key := strconv.FormatInt(videoId, 10)
	exist, err := cache.IsExistComment(s.ctx, key)
	if err != nil {
		return 0, err
	}

	if exist == 1 {
		count, err = cache.CountComments(s.ctx, key)
	} else {
		count, err = db.CountCommentsByVideoID(s.ctx, videoId)
	}

	if err != nil {
		return 0, err
	}
	return
}
