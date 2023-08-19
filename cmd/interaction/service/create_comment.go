package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// CreateComment create comment
func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest) (*db.Comment, error) {

	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, errno.AuthorizationFailedError
	}

	commentModel := &db.Comment{
		VideoId: req.VideoId,
		UserId:  claim.UserId,
		Content: *req.CommentText,
	}
	comment, err := db.CreateComment(s.ctx, commentModel)
	if err != nil {
		return nil, err
	}

	exist, err := cache.IsExistComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	if exist == 1 {
		err = cache.AddComment(s.ctx, req.VideoId,
			&cache.Comment{Id: comment.Id, UserId: comment.UserId, Content: comment.Content},
			float64(comment.CreatedAt.Unix()))
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
