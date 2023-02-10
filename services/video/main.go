package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/go-redis/redis"
	"github.com/ozline/tiktok/pkg/constants"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
	"github.com/ozline/tiktok/services/video/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"
)

var RDB *redis.Client

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", constants.VideoServiceListenAddress)
	svr := video.NewServer(new(TiktokVideoServiceImpl), server.WithServiceAddr(addr))

	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&model.VideoStorageInfo{})

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
