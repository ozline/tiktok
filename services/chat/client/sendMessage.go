package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat/tiktokchatservice"
	"log"
	"time"
)

func main() {
	fmt.Println("----- Chat Client -----")
	userId := 1
	client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
	if err != nil {
		log.Fatal(err)
	}
	//go receive_message_by_id(int64(userId), seqId)
	startTime := time.Now().UnixNano()

	request := &chat.DouyinSendMessageRequest{
		FromUserId: int64(userId),
		ToUserId:   2,
		Content:    "Hello 2,This is 1",
	}

	response, err := client.SendChatMessage(context.Background(), request)

	if err != nil {
		log.Fatal("error", err.Error())
	}
	endTime := time.Now().UnixNano()
	if response.StatusMsg == "Success message" {
		fmt.Println("----- Success To Send Message,Use ", endTime-startTime, " ns -----")
	} else {
		fmt.Println("----- Failure To Send Message -----")
	}

}
