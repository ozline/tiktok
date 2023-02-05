package main

import (
	"context"
	chat "github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
)

// TiktokChatServiceImpl implements the last service interface defined in the IDL.
type TiktokChatServiceImpl struct{}

// SendChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) SendChatMessage(ctx context.Context, req *chat.DouyinMessageChatRequest) (resp *chat.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// AcceptChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) AcceptChatMessage(ctx context.Context, req *chat.Message) (resp *chat.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
