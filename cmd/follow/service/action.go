package service

import (
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/mq"
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
		return errno.UserNotFoundError
	}

	action_meaasge, err := sonic.Marshal(req)
	if err != nil {
		klog.Error(err)
		return err
	}

	err = mq.FollowMQCli.Publish(s.ctx, string(action_meaasge))
	if err != nil {
		klog.Error(err)
		return err
	}

	return nil
}
