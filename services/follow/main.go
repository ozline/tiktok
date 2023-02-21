package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ozline/tiktok/pkg/constants"
	follow "github.com/ozline/tiktok/services/follow/kitex_gen/tiktok/follow/tiktokfollowservice"
	"github.com/ozline/tiktok/services/follow/model"
	"log"
	"net"
)

func init() {
	model.Setup()
}

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", constants.FollowServiceListenAddress)
	svr := follow.NewServer(new(TiktokFollowServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
