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
	fmt.Println("----- getOneVideoInfoTest -----")

	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now().UnixMilli()
	request := &video.GetOneVideoInfoRequest{
		VideoName: "Video5",
	}

	response, err := client.GetOneVideoInfo(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}
	endTime := time.Now().UnixMilli()
	fmt.Println("----- Success To Receive Reponse -----")
	fmt.Println("State=", response.State)
	fmt.Println("VideoId=", response.VideoId)
	fmt.Println("PlayUrl=", response.PlayUrl)
	fmt.Println("CoverUrl=", response.CoverUrl)
	fmt.Println("VideoTitle=", response.VideoTitle)
	fmt.Println("VideoSize=", response.VideoSize)
	fmt.Println("VideoMimeType=", response.VideoMimeType)
	fmt.Println("OwnerName=", response.OwnerName)
	fmt.Println("----- Total Time=", endTime-startTime, "ms -----")
}
