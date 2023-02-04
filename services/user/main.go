package main

import (
	"github.com/ozline/tiktok/services/user/model"
	"log"

	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
)

func main() {
	//初始化数据库
	model.InitDB()
	model.InitRedis()
	svr := tiktokuserservice.NewServer(new(TiktokUserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
