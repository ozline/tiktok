package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/cmd/follow/util"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FollowerList View fan list
func (s *FollowService) FollowerList(req *follow.FollowerListRequest) (*[]*follow.User, error) {
	var userList []*follow.User

	followerList, err := db.FollowerListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, id := range *followerList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		follower := util.ConvertStruct(user) //结构体转换
		userList = append(userList, follower)
	}
	return &userList, nil
}
