package rpc

import (
	"context"

	"github.com/ozline/tiktok/api-gateway/biz/model/model"
	"github.com/ozline/tiktok/kitex_gen/tiktok/follow"
	"github.com/ozline/tiktok/pkg/errno"
)

// RelationAction returns the result of operation T/F
func RelationAction(ctx context.Context, req *follow.RelationActionReq) error {
	resp, err := followClient.RelationAction(ctx, req)

	if err != nil {
		return err
	}

	if resp.StatusCode != errno.SuccessCode {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	return nil
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

// GetFollowList returns the list of users who user follow
func GetFollowList(ctx context.Context, req *follow.UserListReq) (list []*model.User, err error) {
	resp, err := followClient.FollowList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	list = make([]*model.User, 0)
	for _, v := range resp.UserList {
		list = append(list, &model.User{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   *v.FollowCount,
			FollowerCount: *v.FollowCount,
			IsFollow:      v.IsFollow,
		})
	}

	return list, nil
}

// GetFollowerList returns the list of users who follow the user
func GetFollowerList(ctx context.Context, req *follow.UserListReq) (list []*model.User, err error) {
	resp, err := followClient.FollowerList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	list = make([]*model.User, 0)
	for _, v := range resp.UserList {
		list = append(list, &model.User{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   *v.FollowCount,
			FollowerCount: *v.FollowCount,
			IsFollow:      v.IsFollow,
		})
	}

	return list, nil
}

// GetFollowerList returns the list of users who follow the user
func GetFriendList(ctx context.Context, req *follow.UserListReq) (list []*model.User, err error) {
	resp, err := followClient.FriendList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}

	list = make([]*model.User, 0)
	for _, v := range resp.UserList {
		list = append(list, &model.User{
			Id:            v.Id,
			Name:          v.Name,
			FollowCount:   *v.FollowCount,
			FollowerCount: *v.FollowCount,
			IsFollow:      v.IsFollow,
		})
	}

	return list, nil
}
