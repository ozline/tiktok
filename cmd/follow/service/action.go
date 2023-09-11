package service

import (
	"context"
	"sync"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/mq"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

var (
	actionChanGroup = make(map[int64]chan *follow.ActionRequest)
	mu              sync.Mutex
)

// Action Function for the follow/close operation
func (s *FollowService) Action(req *follow.ActionRequest) error {
	// 限流
	if err := cache.Limit(s.ctx, constants.ActionRate, constants.Interval); err != nil {
		return err
	}

	claim, err := utils.CheckToken(req.Token)

	if err != nil {
		return errno.AuthorizationFailedError
	}

	// 禁止自己关注自己
	if claim.UserId == req.ToUserId {
		return errno.FollowYourselfError
	}

	// 判断是否目标用户是否存在
	_, err = rpc.GetUser(s.ctx, &user.InfoRequest{
		UserId: req.ToUserId,
		Token:  req.Token,
	})

	if err != nil {
		return errno.UserNotFoundError
	}

	// 查找是否存在goroutine正在处理该用户的请求
	actionChan, exist := actionChanGroup[claim.UserId]
	if !exist {
		mu.Lock()
		actionChan = make(chan *follow.ActionRequest)
		actionChanGroup[claim.UserId] = actionChan
		klog.Infof("user id : %s, goroutine set up", claim.Id)
		go worker(s.ctx, claim.UserId)
		mu.Unlock()
	}

	// 请求写入通道
	actionChan <- req

	// 写入缓存
	switch req.ActionType {
	case constants.FollowAction:
		// 数据写入redis
		if err := cache.FollowAction(s.ctx, claim.UserId, req.ToUserId); err != nil {
			klog.Error(err)
			return err
		}
	case constants.UnFollowAction:
		time.Sleep(10 * time.Millisecond) // 延迟删除缓存中的数据
		// 删除redis中的数据
		if err = cache.UnFollowAction(s.ctx, claim.UserId, req.ToUserId); err != nil {
			klog.Error(err)
			return err
		}
	default:
		return errno.UnexpectedTypeError
	}

	if err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func worker(ctx context.Context, uid int64) {
	defer func() {
		close(actionChanGroup[uid])
		delete(actionChanGroup, uid)
	}()
	reqSlice := make([]*follow.ActionRequest, 0, 10)

	for {
		select {
		case req := <-actionChanGroup[uid]:
			reqSlice = append(reqSlice, req)
			if len(reqSlice) >= 10 {
				for _, actionReq := range reqSlice {
					action_meaasge, err := sonic.Marshal(actionReq)
					if err != nil {
						klog.Error(err)
					}

					// 消息推送到队列中
					err = mq.FollowMQCli.Publish(ctx, string(action_meaasge))
					if err != nil {
						klog.Error(err)
					}
				}
				return
			}
		case <-time.After(10 * time.Minute):
			for _, actionReq := range reqSlice {
				action_meaasge, err := sonic.Marshal(actionReq)
				if err != nil {
					klog.Error(err)
				}

				// 消息推送到队列中
				err = mq.FollowMQCli.Publish(ctx, string(action_meaasge))
				if err != nil {
					klog.Error(err)
				}
			}
			return
		}
	}
}
