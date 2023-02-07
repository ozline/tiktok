package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	"github.com/ozline/tiktok/pkg/constants"
	auth "github.com/ozline/tiktok/services/auth/kitex_gen/tiktok/auth/tiktokauthservice"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", constants.AuthServiceListenAddress)
	svr := auth.NewServer(new(TiktokAuthServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
