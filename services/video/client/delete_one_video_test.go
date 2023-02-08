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

	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8892"))
	if err != nil {
		log.Fatal(err)
	}

	request := &video.DeleteVideoRequest{
		Title:       "Video5",
		DeletorName: "sunyiwen",
	}
	startTime := time.Now().UnixMilli()
	response, err := client.DeleteVideo(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}
	endTime := time.Now().UnixMilli()
	fmt.Println("----- Success To Receive Reponse -----")
	fmt.Println("State=", response.State)
	fmt.Println("DeleteVideoName=", response.DeleteVideoName)
	fmt.Println("DeletorName=", response.DeletorName)
	fmt.Println("VideoOwnerName=", response.VideoOwnerName)
	fmt.Println("----- Total Time=", endTime-startTime, "ms -----")
}
