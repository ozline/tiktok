package rpc

import (
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/kitex_gen/follow/followservice"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
)

var (
	followClient      followservice.Client
	interactionClient interactionservice.Client
	videoClient       videoservice.Client
)

func Init() {
	InitFollowRPC()
	InitInteractionRPC()
	InitVideoRPC()
}
