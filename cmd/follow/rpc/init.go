package rpc

import (
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
	chatClient messageservice.Client
)

func Init() {
	InitUserRPC()
	InitChatRPC()
}
