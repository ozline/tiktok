package service

import (
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FollowList View the follow list
func (s *FollowService) FollowList(req *follow.FollowListRequest) (*[]*follow.User, error) {
	// 限流
	if err := cache.Limit(s.ctx); err != nil {
		return nil, err
	}

	userList := make([]*follow.User, 0, 10)

	// 先查redis
	followList, err := cache.FollowListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(*followList) == 0 { // redis中查不到再查db
		followList, err = db.FollowListAction(s.ctx, req.UserId)
		if errors.Is(err, db.RecordNotFound) { // db中也查不到
			klog.Info("you are not following anyone")
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		// db中查到后写入redis
		err := cache.UpdateFollowList(s.ctx, req.UserId, followList)
		if err != nil {
			return nil, err
		}
	}

	// 数据处理
	for _, id := range *followList {
		user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: id,
			Token:  req.Token,
		})
		if err != nil {
			return nil, err
		}
		follow := pack.User(user) // 结构体转换
		userList = append(userList, follow)
	}

	return &userList, nil
}
