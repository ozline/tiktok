package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// Action Function for the follow/close operation
func (s *FollowService) Action(req *follow.ActionRequest) error {
	//限流
	if err := cache.Limit(s.ctx); err != nil {
		return err
	}

	claim, err := utils.CheckToken(req.Token)

	if err != nil {
		return errno.AuthorizationFailedError
	}

	action := &db.Follow{
		UserId:   claim.UserId,
		ToUserId: req.ToUserId,
	}

	if req.ActionType == 1 {
		//数据写入redis
		if err := cache.FollowAction(s.ctx, action.UserId, action.ToUserId); err != nil {
			return err
		}
		//数据写入db/更改db数据
		if err = db.FollowAction(s.ctx, action); err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		//删除redis中的数据
		if err = cache.UnFollowAction(s.ctx, action.UserId, action.ToUserId); err != nil {
			return err
		}
		//更改db数据
		if err = db.UnFollowAction(s.ctx, action); err != nil {
			return err
		}
	} else {
		return errno.UnexpectedTypeError
	}

	if err != nil {
		return err
	}
	return nil
}
