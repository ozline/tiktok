package main

import (
	chat "github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"log"
)

func main() {
	svr := chat.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
