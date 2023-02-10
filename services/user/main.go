package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/user/configs"
	user "github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"log"
	"net"
)

func main() {
	//初始化数据库
	configs.InitDB()

	addr, _ := net.ResolveTCPAddr("tcp", constants.UserServiceListenAddress)
	svr := user.NewServer(new(TiktokUserServiceImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
