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
	fmt.Println("----- downMultiVideoTest -----")
	startTime := time.Now().UnixMilli()
	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	request := &video.DownloadOneVideoRequest{
		VideoName: "Video5",
	}

	response, err := client.DownloadOneVideo(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}
	endTime := time.Now().UnixMilli()
	fmt.Println("----- Success To Receive Reponse -----")
	fmt.Println("State=", response.State)
	fmt.Println("VideoTitle=", response.VideoTitle)
	fmt.Println("VideoUrl=", response.VideoUrl)
	fmt.Println("OwnerName=", response.OwnerName)
	fmt.Println("----- Total Time=", endTime-startTime, "ms -----")
}
