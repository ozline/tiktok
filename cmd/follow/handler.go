package main

import (
	"context"

	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/service"
	follow "github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Action implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Action(ctx context.Context, req *follow.ActionRequest) (resp *follow.ActionResponse, err error) {
	resp = new(follow.ActionResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	if err := service.NewFollowService(ctx).Action(req); err != nil {
		resp.Base = pack.BuildBaseResp(err)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListRequest) (resp *follow.FollowListResponse, err error) {
	resp = new(follow.FollowListResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	followResp, err := service.NewFollowService(ctx).FollowList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserList = *followResp
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (resp *follow.FollowerListResponse, err error) {
	resp = new(follow.FollowerListResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	followerResp, err := service.NewFollowService(ctx).FollowerList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserList = *followerResp
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	resp = new(follow.FriendListResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	friendResp, err := service.NewFollowService(ctx).FriendList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.UserList = *friendResp
	return
}
