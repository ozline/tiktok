package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FriendList Viewing friends list
func (s *FollowService) FriendList(req *follow.FriendListRequest) (*[]user.User, error) {
	var userList []user.User

	friendList, err := db.FriendListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, id := range *friendList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		userList = append(userList, *user)
	}
	return &userList, nil
}
