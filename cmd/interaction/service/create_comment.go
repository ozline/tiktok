package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
	"strconv"
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

	key := strconv.FormatInt(comment.VideoId, 10)
	exist, err := cache.IsExistComment(s.ctx, key)
	if err != nil {
		return nil, err
	}

	if exist == 1 {
		err = cache.AddComment(s.ctx, key, comment)
		if err != nil {
			return nil, err
		}
	}

	return comment, nil
}
