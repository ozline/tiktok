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
	//fmt.Println("----- downOneVideoTest -----")

	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8892"))
	if err != nil {
		log.Fatal(err)
	}

	startTime := time.Now().UnixMilli()
	request := &video.DownloadOneVideoRequest{
		VideoName: "Video3",
	}

	response, err := client.DownloadOneVideo(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}
	endTime := time.Now().UnixMilli()
	//fmt.Println("----- Success To Receive Reponse -----")
	//fmt.Println("State=", response.State)
	//fmt.Println("VideoTitle=", response.VideoTitle)
	//fmt.Println("VideoUrl=", response.VideoUrl)
	//fmt.Println("OwnerName=", response.OwnerName)
	if response.State == true {
		fmt.Println("----- DownOneVideoTest :", endTime-startTime, "ms -----")
	} else {
		fmt.Println("----- DownOneVideoTest Failed -----")
	}

}
