package rpc

import (
	"github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
)

var (
	userClient  tiktokuserservice.Client
	authClient  tiktokauthservice.Client
	videoClient tiktokvideoservice.Client
)

func Init() {
	initAuthRPC()
	initUserRPC()
	InitVideoRPC()
}
