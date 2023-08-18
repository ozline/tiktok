package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
	"strconv"
)

// CreateComment create comment
func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest) (*db.Comment, error) {

	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}

	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, errno.AuthorizationFailedError
	}

	commentModel := &db.Comment{
		VideoId: videoId,
		UserId:  claim.UserId,
		Content: *req.CommentText,
	}

	return db.CreateComment(s.ctx, commentModel)
}
