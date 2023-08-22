package rpc

import (
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	interactionClient interactionservice.Client
)

func Init() {
	InitUserRPC()
	InitInteractionRPC()
}
