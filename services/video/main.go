package main

import (
	"log"

	video "github.com/ozline/tiktok/kitex_gen/tiktok/video/tiktokvideoservice"

	"github.com/ozline/tiktok/services/video/dal"
)

func Init() {
	dal.Init()
}

func main() {
	Init()
	svr := video.NewServer(new(TiktokVideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
