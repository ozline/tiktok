package rpc

import "github.com/ozline/tiktok/kitex_gen/user/userservice"

var (
	userClient userservice.Client
)

func Init() {
	InitUserRPC()
	InitChatRPC()
}
