package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/cmd/follow/util"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FriendList Viewing friends list
func (s *FollowService) FriendList(req *follow.FriendListRequest) (*[]*follow.FriendUser, error) {
	var friendList []*follow.FriendUser

	userList, err := db.FriendListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	for _, id := range *userList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		friend := util.ConvertStruct(user) //结构体转换

		message, msgType, err := rpc.GetMessage(s.ctx, &chat.MessageListRequest{
			Token:    req.Token,
			ToUserId: req.UserId,
		}, req.UserId, id)

		if err != nil {
			return nil, err
		}

		friendUser := &follow.FriendUser{
			User:    friend,
			Message: &message,
			MsgType: msgType,
		}
		friendList = append(friendList, friendUser)

	}
	return &friendList, nil
}
