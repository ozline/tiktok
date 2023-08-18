package service

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) (*[]db.Comment, error) {
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}

	return db.GetCommentsByVideoID(s.ctx, videoId)
}
