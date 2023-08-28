package service

import (
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// Action Function for the follow/close operation
func (s *FollowService) Action(req *follow.ActionRequest) error {
	// 限流
	if err := cache.Limit(s.ctx, constants.ActionRate, constants.Interval); err != nil {
		return err
	}

	claim, err := utils.CheckToken(req.Token)

	if err != nil {
		return errno.AuthorizationFailedError
	}

	// 禁止自己关注自己
	if claim.UserId == req.ToUserId {
		return errno.FollowYourselfError
	}

	// 判断是否目标用户是否存在
	_, err = rpc.GetUser(s.ctx, &user.InfoRequest{
		UserId: req.ToUserId,
		Token:  req.Token,
	})

	if err != nil {
		klog.Info(err)
		return errno.UserNotFoundError
	}

	action := &db.Follow{
		UserID:   claim.UserId,
		ToUserID: req.ToUserId,
	}

	switch req.ActionType {
	case constants.FollowAction:
		// 数据写入redis
		if err := cache.FollowAction(s.ctx, action.UserID, action.ToUserID); err != nil {
			return err
		}
		// 数据写入db/更改db数据
		if err = db.FollowAction(s.ctx, action); err != nil {
			return err
		}
	case constants.UnFollowAction:
		// 更改db数据
		if err = db.UnFollowAction(s.ctx, action); err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond) // 延迟删除缓存中的数据
		// 删除redis中的数据
		if err = cache.UnFollowAction(s.ctx, action.UserID, action.ToUserID); err != nil {
			return err
		}
	default:
		return errno.UnexpectedTypeError
	}

	if err != nil {
		return err
	}

	return nil
}
