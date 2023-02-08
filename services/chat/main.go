package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ozline/tiktok/pkg/constants"
	chat "github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat/tiktokchatservice"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", constants.ChatServiceListenAddress)
	svr := chat.NewServer(new(TiktokChatServiceImpl), server.WithServiceAddr(addr))

	dbr, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	dbr.AutoMigrate(&Message{})

	dbs, err := gorm.Open(sqlite.Open("sendMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	dbs.AutoMigrate(&Message{})

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
