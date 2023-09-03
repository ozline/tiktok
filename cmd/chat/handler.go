package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	pack "github.com/ozline/tiktok/cmd/chat/pack"
	service "github.com/ozline/tiktok/cmd/chat/service"
	chat "github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessagePost implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessagePost(ctx context.Context, req *chat.MessagePostRequest) (resp *chat.MessagePostReponse, err error) {
	resp = new(chat.MessagePostReponse)
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, err
	}

	err = service.NewChatService(ctx).SendMessage(req, claim.UserId, time.Now().Format(time.RFC3339))
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, err
	}
	resp.Base = pack.BuildBaseResp(nil)
	return
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *chat.MessageListRequest) (resp *chat.MessageListResponse, err error) {
	resp = new(chat.MessageListResponse)
	claim, err := utils.CheckToken(req.Token)
	if err != nil || claim == nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, err
	}
	// 获取消息列表

	// redis->mysql
	// redis中存在则返回，不存在查询mysql,
	messageList, err := service.NewChatService(ctx).GetMessages(req, claim.UserId)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		resp.MessageList = pack.BuildMessage(nil)
		resp.Total = 0
		return resp, err
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.MessageList = pack.BuildMessage(messageList)
	resp.Total = int64(len(messageList))
	return
}
