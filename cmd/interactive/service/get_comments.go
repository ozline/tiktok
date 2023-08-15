package service

import (
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"strconv"
)

func (s *CommentService) GetComments(req *interactive.CommentListRequest) (*[]db.Comment, error) {
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}

	return db.GetCommentsByVideoID(s.ctx, videoId)
}
