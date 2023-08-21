package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"strconv"
)

// CreateComment create comment
func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest, userId int64) (*db.Comment, error) {

	commentModel := &db.Comment{
		VideoId: req.VideoId,
		UserId:  userId,
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
