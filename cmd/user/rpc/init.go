package rpc

import (
	"github.com/ozline/tiktok/kitex_gen/follow/followservice"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/kitex_gen/video/videoservice"
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
