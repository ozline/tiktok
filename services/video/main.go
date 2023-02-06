package main

import (
	"github.com/go-redis/redis"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var RDB *redis.Client

func main() {
	svr := video.NewServer(new(TiktokVideoServiceImpl))
	db, err := gorm.Open(sqlite.Open("videoStorage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoStorageInfo{})

	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	go PerioUpdateVideoCache(10)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
