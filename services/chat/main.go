package main

import (
	chat "github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat/tiktokchatservice"
	"github.com/ozline/tiktok/services/chat/utils"
	"log"
)

func main() {
	svr := chat.NewServer(new(TiktokChatServiceImpl))
	//开启配置
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
