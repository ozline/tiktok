package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"

	"strconv"
)

// DeleteComment delete comment
func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest) (*db.Comment, error) {

	commentId, err := strconv.ParseInt(*req.CommentId, 10, 64)
	if err != nil {
		return nil, err
	}

	comment, err := db.GetCommentByID(s.ctx, commentId)
	if err != nil {
		return nil, err
	}

	comment, err = db.DeleteComment(s.ctx, comment)
	if err != nil {
		return nil, err
	}

	exist, err := cache.IsExistComment(s.ctx, comment.VideoId)
	if err != nil {
		return nil, err
	}
	if exist == 1 {
		err = cache.DeleteComment(s.ctx, comment.VideoId, comment)
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
