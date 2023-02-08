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

	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8892"))
	if err != nil {
		log.Fatal(err)
	}
	startTime := time.Now().UnixMilli()
	number := 5
	request := &video.DownloadMultiVideoRequest{
		VideoNumber:     int64(number),
		Downloador_Name: "sunyiwen",
	}
	response, err := client.DownloadMultiVideo(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}

	endTime := time.Now().UnixMilli()
	fmt.Println("----- Success To Receive Reponse -----")
	fmt.Println("State=", response.State)
	fmt.Println("VideoNumber=", response.VideoNumber)
	fmt.Println("--------------------------------------")
	for i := 1; i <= number; i++ {
		//fmt.Println("Title", i, "=", response.VideoTitles[i-1])
		fmt.Println("Url", i, "=", response.VideoUrls[i-1])
		//fmt.Println("Owner", i, "=", response.OwnerNames[i-1])
	}
	fmt.Println("----- Total Time=", endTime-startTime, "ms -----")
}
