package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FollowerList View fan list
func (s *FollowService) FollowerList(req *follow.FollowerListRequest) (*[]*follow.User, error) {
	//限流
	if err := cache.Limit(s.ctx); err != nil {
		return nil, err
	}

	userList := make([]*follow.User, 0, 10)

	//先查redis
	followerList, err := cache.FollowerListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(*followerList) == 0 { //redis中查不到再查db
		followerList, err = db.FollowerListAction(s.ctx, req.UserId)
		if errors.Is(err, db.RecordNotFound) { //db中也查不到
			return nil, errors.New("you do not have any followers")
		} else if err != nil {
			return nil, err
		}
		//db中查到后写入redis
		err := cache.UpdateFollowerList(s.ctx, req.UserId, followerList)
		if err != nil {
			return nil, err
		}
	}

	//数据处理
	for _, id := range *followerList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		follower := pack.User(user) //结构体转换
		userList = append(userList, follower)
	}
	return &userList, nil
}
