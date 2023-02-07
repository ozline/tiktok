package main

import (
	"github.com/ozline/tiktok/services/user/configs"
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	"github.com/ozline/tiktok/pkg/constants"
	user "github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
)

func main() {

	configs.InitDB()

	addr, _ := net.ResolveTCPAddr("tcp", constants.UserServiceListenAddress)
	svr := user.NewServer(new(TiktokUserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
