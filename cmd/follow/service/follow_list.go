package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FollowList View the follow list
func (s *FollowService) FollowList(req *follow.FollowListRequest) (*[]user.User, error) {
	var userList []user.User

	followList, err := db.FollowListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, id := range *followList {
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
