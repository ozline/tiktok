package rpc

import (
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"github.com/ozline/tiktok/kitex_gen/follow/followservice"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	followClient      followservice.Client
	interactionClient interactionservice.Client
	chatClient        messageservice.Client
	videoClient       videoservice.Client
)

func Init() {
	InitUserRPC()
	InitFollowRPC()
	InitInteractionRPC()
	InitChatRPC()
	InitVideoRPC()
}
