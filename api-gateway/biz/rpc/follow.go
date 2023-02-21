package rpc

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/follow"
	"github.com/ozline/tiktok/pkg/errno"
)

// RelationAction returns the result of operation T/F
func RelationAction(ctx context.Context, req *follow.RelationActionReq) (bool, error) {
	resp, err := followClient.RelationAction(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return false, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	return true, nil
}

// RelationQuery returns the relation between two users: 1 = NoFollow, 2 = Following, 3 = Followed, 4 = Friend
func RelationQuery(ctx context.Context, req *follow.RelationQueryReq) (int64, error) {
	resp, err := followClient.RelationQuery(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	return int64(resp.RelationCode), nil
}

// GetFollowList returns the list of users who follow the user
func GetFollowList(ctx context.Context, req *follow.UserListReq) (int64, error) {
	resp, err := followClient.FollowList(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	return 0, nil
}
