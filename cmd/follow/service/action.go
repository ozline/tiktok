package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// Action Function for the follow/close operation
func (s *FollowService) Action(req *follow.ActionRequest) error {
	claim, err := utils.CheckToken(req.Token)

	if err != nil {
		return errno.AuthorizationFailedError
	}

	action := &db.Follow{
		UserId:     claim.UserId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	}

	return db.FollowAction(s.ctx, action)
}
