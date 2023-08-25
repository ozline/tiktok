package service

import (
	"errors"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// FollowerList View fan list
func (s *FollowService) FollowerList(req *follow.FollowerListRequest) (*[]*follow.User, error) {
	// 限流
	if err := cache.Limit(s.ctx, 100, 1*time.Second); err != nil {
		return nil, err
	}

	userList := make([]*follow.User, 0, 10)
	var wg sync.WaitGroup
	var mu sync.Mutex
	isErr := false

	// 先查redis
	followerList, err := cache.FollowerListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(*followerList) == 0 { // redis中查不到再查db
		followerList, err = db.FollowerListAction(s.ctx, req.UserId)
		if errors.Is(err, db.RecordNotFound) { // db中也查不到
			klog.Info("you do not have any followers")
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		// db中查到后写入redis
		err := cache.UpdateFollowerList(s.ctx, req.UserId, followerList)
		if err != nil {
			return nil, err
		}
	}

	// 数据处理
	for _, id := range *followerList {
		wg.Add(1)
		go func(id int64, req *follow.FollowerListRequest, userList *[]*follow.User, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()

			user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: id,
				Token:  req.Token,
			})
			if err != nil {
				mu.Lock()
				isErr = true // 报错就修改为true
				mu.Unlock()
				return
			}

			follow := pack.User(user) // 结构体转换

			mu.Lock()
			*userList = append(*userList, follow)
			mu.Unlock()
		}(id, req, &userList, &wg, &mu)
		if isErr {
			return nil, errors.New("RPC call error")
		}
	}

	wg.Wait()

	return &userList, nil
}
