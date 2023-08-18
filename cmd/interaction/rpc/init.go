package rpc

import (
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var (
	userClient  userservice.Client
	videoClient videoservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
}
