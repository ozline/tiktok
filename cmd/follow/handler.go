package main

import (
	"context"

	follow "github.com/ozline/tiktok/kitex_gen/follow"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Action implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Action(ctx context.Context, req *follow.ActionRequest) (resp *follow.ActionResponse, err error) {
	// TODO: Your code here...
	resp = new(follow.ActionResponse)
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListRequest) (resp *follow.FollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(follow.FollowListResponse)
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (resp *follow.FollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(follow.FollowerListResponse)
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(follow.FriendListResponse)
	return
}
