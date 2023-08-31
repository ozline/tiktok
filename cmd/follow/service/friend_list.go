package service

import (
	"errors"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/cmd/follow/pack"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
)

// FriendList Viewing friends list
func (s *FollowService) FriendList(req *follow.FriendListRequest) (*[]*follow.FriendUser, error) {
	// 限流
	if err := cache.Limit(s.ctx, constants.FriendListRate, constants.Interval); err != nil {
		return nil, err
	}

	friendList := make([]*follow.FriendUser, 0, 10)
	var wg sync.WaitGroup
	var mu sync.Mutex
	isErr := false

	// 先查redis
	userList, err := cache.FriendListAction(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(*userList) == 0 { // redis中查不到再查db
		userList, err = db.FriendListAction(s.ctx, req.UserId)
		if errors.Is(err, db.RecordNotFound) { // db中也查不到
			return &friendList, nil
		} else if err != nil {
			return nil, err
		}
		// db中查到后写入redis
		followList, _ := db.FollowListAction(s.ctx, req.UserId)
		followerList, _ := db.FollowerListAction(s.ctx, req.UserId)
		err := cache.UpdateFriendList(s.ctx, req.UserId, followList, followerList)
		if err != nil {
			return nil, err
		}
	}

	// 数据处理
	for _, userID := range *userList {
		wg.Add(1)
		go func(id int64, req *follow.FriendListRequest, userList *[]*follow.FriendUser, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer func() {
				// 协程内部使用recover捕获可能在调用逻辑中发生的panic
				if e := recover(); e != nil {
					// 某个服务调用协程报错，在这里打印一些错误日志
					klog.Info("recover panic:", e)
				}
				wg.Done()
			}()

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
			friend := pack.User(user) // 结构体转换

			message, msgType, err := rpc.GetMessage(s.ctx, &chat.MessageListRequest{
				Token:    req.Token,
				ToUserId: req.UserId,
			}, req.UserId, id)
			if err != nil {
				mu.Lock()
				isErr = true // 报错就修改为true
				mu.Unlock()
				return
			}

			friendUser := &follow.FriendUser{
				User:    friend,
				Message: &message,
				MsgType: msgType,
			}

			mu.Lock()
			*userList = append(*userList, friendUser)
			mu.Unlock()
		}(userID, req, &friendList, &wg, &mu)
	}

	wg.Wait()

	if isErr {
		return nil, errno.ServiceError
	}

	return &friendList, nil
}
