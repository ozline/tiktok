package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"strconv"
)

// DeleteComment delete comment
func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest, userId int64) (*db.Comment, error) {

	comment, err := db.GetCommentByID(s.ctx, *req.CommentId)
	if err != nil {
		return nil, err
	}

	if comment.UserId != userId {
		return nil, errno.AuthorizationFailedError
	}

	comment, err = db.DeleteComment(s.ctx, comment)
	if err != nil {
		return nil, err
	}

	key := strconv.FormatInt(comment.VideoId, 10)
	exist, err := cache.IsExistComment(s.ctx, key)
	if err != nil {
		return nil, err
	}
	if exist == 1 {
		err = cache.DeleteComment(s.ctx, key, comment)
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
