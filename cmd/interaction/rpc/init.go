package rpc

import (
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
	"github.com/ozline/tiktok/kitex_gen/video/videoservice"
)

var (
	userClient  userservice.Client
	videoClient videoservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
}
