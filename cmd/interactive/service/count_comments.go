package service

import (
	"github.com/ozline/tiktok/cmd/interactive/dal/cache"
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"strconv"
)

func (s *CommentService) CountComments(req *interactive.CommentCountRequest) (count int64, err error) {
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return 0, err
	}

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
