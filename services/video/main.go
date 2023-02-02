package main

import (
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {

	svr := video.NewServer(new(TiktokVideoServiceImpl))
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&VideoInfo{})

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
