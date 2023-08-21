package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FriendList Viewing friends list
func (s *FollowService) FriendList(req *follow.FriendListRequest) (*[]*follow.FriendUser, error) {
	//限流
	if err := cache.Limit(s.ctx); err != nil {
		return nil, err
	}

	var friendList []*follow.FriendUser

	//先查redis
	userList, err := cache.FriendListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(*userList) == 0 { //redis中查不到再查db
		userList, err = db.FriendListAction(s.ctx, req.UserId)
		if errors.Is(err, db.RecordNotFound) { //db中也查不到
			return nil, errors.New("you do not have any friends")
		} else if err != nil {
			return nil, err
		}
		//db中查到后写入redis
		followList, _ := db.FollowListAction(s.ctx, req.UserId)
		followerList, _ := db.FollowerListAction(s.ctx, req.UserId)
		err := cache.UpdateFriendList(s.ctx, req.UserId, followList, followerList)
		if err != nil {
			return nil, err
		}
	}

	//数据处理
	for _, id := range *userList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		friend := pack.User(user) //结构体转换

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
