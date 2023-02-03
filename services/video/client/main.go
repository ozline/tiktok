package main

import (
	"context"
	"log"

	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := tiktokvideoservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req1 := &video.Request1{
		Message: "test video PingPong",
	}
	resp, err := client.PingPong(context.Background(), req1)
	if err != nil {
		log.Fatal("err1", err.Error())
	}
	log.Println("PingPong Func Response", resp)
}
