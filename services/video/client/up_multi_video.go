package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
	"log"
	"time"
)

func main() {
	//fmt.Println("----- UpMultiVideoTest -----")

	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8892"))
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now().UnixMilli()
	for i := 1; i <= 30; i++ {
		picstr := fmt.Sprintf("%s%d%s", "/home/ubuntu/Desktop/QingXuYing/tiktok/services/video/client/picture/pic", i, ".jpg")
		vidstr := fmt.Sprintf("%s%d%s", "/home/ubuntu/Desktop/QingXuYing/tiktok/services/video/client/video/video", i, ".mp4")
		title := fmt.Sprintf("%s%d", "Video", i)

		request := &video.PutVideoRequest{
			PlayUrl:   vidstr,
			CoverUrl:  picstr,
			Title:     title,
			OwnerName: "sunyiwen",
		}
		_, err := client.PutVideo(context.Background(), request)
		if err != nil {
			log.Fatal("err1", err.Error())
		}

	}
	endTime := time.Now().UnixMilli()
	fmt.Println("----- UpMultiVideoTest :", endTime-startTime, "ms -----")
}
