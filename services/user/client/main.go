package main

import (
	"context"
	"log"

	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := tiktokuserservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req1 := &user.Request1{
		Message: "testPingPong",
	}
	resp, err := client.PingPong(context.Background(), req1)
	if err != nil {
		log.Fatal("err1", err.Error())
	}
	log.Println("PingPong Func Response", resp)
}
