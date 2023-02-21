package main

import (
	"context"

	follow "github.com/ozline/tiktok/kitex_gen/tiktok/follow"
	"github.com/ozline/tiktok/services/follow/service"
)

// TiktokFollowServiceImpl implements the last service interface defined in the IDL.
type TiktokFollowServiceImpl struct{}

// Ping implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) Ping(ctx context.Context, req *follow.PingReq) (resp *follow.BaseRsp, err error) {
	resp = new(follow.BaseRsp)
	resp.StatusCode = service.Success
	resp.StatusMsg = service.GetMsg(resp.StatusCode)
	return
}

// RelationAction implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) RelationAction(ctx context.Context, req *follow.RelationActionReq) (resp *follow.BaseRsp, err error) {
	resp = new(follow.BaseRsp)
	switch req.ActionType {
	// ActionType 1, 关注
	case 1:
		resp.StatusCode = service.AddFollowRelation(req.UserId, req.ToUserId)
		resp.StatusMsg = service.GetMsg(resp.StatusCode)
	// ActionType 2, 取关
	case 2:
		resp.StatusCode = service.RemoveFollowRelation(req.UserId, req.ToUserId)
		resp.StatusMsg = service.GetMsg(resp.StatusCode)
	// unDefinition Action
	default:
		resp.StatusCode = service.InvalidRelationType
		resp.StatusMsg = service.GetMsg(resp.StatusCode)
	}
	return
}

// FollowList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FollowList(ctx context.Context, req *follow.UserListReq) (resp *follow.UserListRsp, err error) {
	resp = new(follow.UserListRsp)
	resp.UserList, resp.StatusCode = service.ListFollows(req.UserId, req.PageNum, req.PageSize)
	resp.StatusMsg = service.GetMsg(resp.StatusCode)
	return
}

// FollowerList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FollowerList(ctx context.Context, req *follow.UserListReq) (resp *follow.UserListRsp, err error) {
	resp = new(follow.UserListRsp)
	resp.UserList, resp.StatusCode = service.ListFollowers(req.UserId, req.PageNum, req.PageSize)
	resp.StatusMsg = service.GetMsg(resp.StatusCode)
	return
}

// FriendList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FriendList(ctx context.Context, req *follow.UserListReq) (resp *follow.UserListRsp, err error) {
	resp = new(follow.UserListRsp)
	resp.UserList, resp.StatusCode = service.ListFriends(req.UserId, req.PageNum, req.PageSize)
	resp.StatusMsg = service.GetMsg(resp.StatusCode)
	return
}

// RelationQuery implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) RelationQuery(ctx context.Context, req *follow.RelationQueryReq) (resp *follow.RelationQueryRsp, err error) {
	resp = new(follow.RelationQueryRsp)
	resp.RelationCode = service.QueryRelation(req.UserId, req.ToUserId)
	resp.StatusCode = service.Success
	resp.StatusMsg = service.GetMsg(resp.StatusCode)
	return
}
