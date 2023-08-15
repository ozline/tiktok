package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/cmd/follow/util"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FriendList Viewing friends list
func (s *FollowService) FriendList(req *follow.FriendListRequest) (*[]*follow.User, error) {
	var userList []*follow.User

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
		friend := util.ConvertStruct(user) //结构体转换
		//TODO:转为FriendUser
		userList = append(userList, friend)
	}
	return &userList, nil
}
