package service

import "context"

type ChatService struct{
    ctx context.Context   
}

// NewUserService new UserService
func NewChatService(ctx context.Context) *ChatService {
	return &ChatService{ctx: ctx}
}