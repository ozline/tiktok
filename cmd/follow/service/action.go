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
		UserId:   claim.UserId,
		ToUserId: req.ToUserId,
	}

	if req.ActionType == 1 {
		err = db.FollowAction(s.ctx, action)
	} else if req.ActionType == 2 {
		err = db.UnFollowAction(s.ctx, action)
	} else {
		return errno.UnexpectedTypeError
	}

	if err != nil {
		return err
	}

	return nil
}
