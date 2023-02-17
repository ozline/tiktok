package main

import (
	"context"
	"github.com/ozline/tiktok/services/follow/model"
	"github.com/ozline/tiktok/services/follow/service"

	"github.com/ozline/tiktok/services/follow/kitex_gen/tiktok/follow"
)

// TiktokFollowServiceImpl implements the last service interface defined in the IDL.
type TiktokFollowServiceImpl struct{}

// Ping implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) Ping(ctx context.Context, req *follow.PingReq) (resp *follow.PingRsp, err error) {
	resp = new(follow.PingRsp)
	resp.Message = "Pong"
	return resp, nil
}

// RelationAction implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) RelationAction(ctx context.Context, req *follow.RelationActionReq) (resp *follow.RelationActionRsp, err error) {
	resp = new(follow.RelationActionRsp)

	switch req.ActionType {
	// ActionType 1 关注用户
	case 1:
		// FIXME 传进来 token 谁去鉴权啊草
		if statusCode := model.AddRelation(req., req.ToUserId); statusCode != 0 {

		} else {
			resp.StatusCode = 0
			var msg = "Success"
			resp.StatusMsg = &msg
			return resp, nil
		}
	// ActionType 2 取消关注
	case 2:
		// FIXME 传进来 token 谁去鉴权啊草
		if statusCode := model.RemoveRelation(req.Token, req.ToUserId); statusCode != 0 {

		} else {
			resp.StatusCode = 0
			var msg = "Success"
			resp.StatusMsg = &msg
			return resp, nil
		}
	}

	return
}

// FollowList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListReq) (resp *follow.FollowListRsp, err error) {

	service.QueryFollowList(req.UserId)

	return
}

// FollowerList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListReq) (resp *follow.FollowerListRsp, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the TiktokFollowServiceImpl interface.
func (s *TiktokFollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListReq) (resp *follow.FriendListRsp, err error) {
	// TODO: Your code here...
	return
}
