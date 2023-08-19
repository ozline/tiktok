package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"

)

// DeleteComment delete comment
func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest) (*db.Comment, error) {

	comment, err := db.GetCommentByID(s.ctx, *req.CommentId)
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
		err = cache.DeleteComment(s.ctx, comment.VideoId,
			&cache.Comment{Id: comment.Id, UserId: comment.UserId, Content: comment.Content})
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
