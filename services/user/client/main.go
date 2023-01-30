package main

import (
	"context"
	"log"

	"github.com/ozline/tiktok/services/user/kitex_gen/kitex_gen"
	"github.com/ozline/tiktok/services/user/kitex_gen/kitex_gen/kitexprotobuf"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := kitexprotobuf.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req1 := &kitex_gen.Request1{
		Message: "message1",
	}
	resp, err := client.MyHandT1(context.Background(), req1)
	if err != nil {
		log.Fatal("err1", err.Error())
	}
	log.Println("MyHandT1 Func Response", resp)
}
