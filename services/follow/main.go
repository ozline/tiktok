package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	follow "github.com/ozline/tiktok/kitex_gen/tiktok/follow/tiktokfollowservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/follow/model"
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
