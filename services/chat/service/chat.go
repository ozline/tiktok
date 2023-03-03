package service

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/services/chat/dal/db"
)

type ChatService struct {
	ctx context.Context
}

func NewChatService(ctx context.Context) *ChatService {
	return &ChatService{ctx: ctx}
}

func (cs *ChatService) SendMessage(req *chat.SendMessageRequest) (*chat.ChatMsg, error) {
	return db.CreateMessage(cs.ctx, req)
}

func (cs *ChatService) GetChatList(req *chat.ReceiveMessageRequest) ([]*chat.ChatMsg, error) {
	return db.GetMessageList(cs.ctx, req)
}
