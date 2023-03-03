package main

import (
	"context"

	chat "github.com/ozline/tiktok/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/chat/pack"
	"github.com/ozline/tiktok/services/chat/service"
)

// TiktokChatServiceImpl implements the last service interface defined in the IDL.
type TiktokChatServiceImpl struct{}

// SendChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) SendChatMessage(ctx context.Context, req *chat.SendMessageRequest) (resp *chat.SendMessageResponse, err error) {

	resp = new(chat.SendMessageResponse)

	msg, err := service.NewChatService(ctx).SendMessage(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = msg
	return resp, nil
}

// AcceptChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) AcceptChatMessage(ctx context.Context, req *chat.ReceiveMessageRequest) (resp *chat.ReceiveMessageResponse, err error) {

	resp = new(chat.ReceiveMessageResponse)

	msg, err := service.NewChatService(ctx).GetChatList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Data = msg

	return resp, nil
}
