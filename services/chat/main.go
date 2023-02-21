package main

import (
	"log"
	"net"

	"github.com/cloudwego/kitex/server"
	chat "github.com/ozline/tiktok/kitex_gen/tiktok/chat/tiktokchatservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/chat/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", constants.ChatServiceListenAddress)
	svr := chat.NewServer(new(TiktokChatServiceImpl), server.WithServiceAddr(addr))

	dbr, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	dbr.AutoMigrate(&model.Message{})

	dbs, err := gorm.Open(sqlite.Open("sendMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	dbs.AutoMigrate(&model.Message{})

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
