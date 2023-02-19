package main

import (
	"context"
	"github.com/cloudwego/kitex/server"
	"github.com/go-redis/redis"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
	"github.com/ozline/tiktok/services/video/model"
	"github.com/ozline/tiktok/services/video/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"
)

var RDB *redis.Client

func main() {

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8892")
	svr := video.NewServer(new(TiktokVideoServiceImpl), server.WithServiceAddr(addr))

	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&model.VideoStorageInfo{})

	service.NewTokenService(context.Background()).TokenLimits()
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
