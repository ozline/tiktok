package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user"
	"github.com/ozline/tiktok/services/user/kitex_gen/tiktok/user/tiktokuserservice"
	"log"
)

func main() {
	//测试
	client, err := tiktokuserservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	//--------------------运行测试-------------------------
	//req1 := &user.Request1{
	//	Message: "testPingPong",
	//}
	//resp, err := client.PingPong(context.Background(), req1)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//log.Println("PingPong Func Response", resp)
	//--------------------注册测试-------------------------
	registerRequest := &user.DouyinUserRegisterRequest{}

	registerResponse, err := client.Register(context.Background(), registerRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(registerResponse)
}
