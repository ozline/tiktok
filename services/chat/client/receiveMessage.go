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
	userId := 1
	receive_message_by_id(int64(userId))
}

func receive_message_by_id(userId int64) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
		if err != nil {
			log.Fatal(err)
		}
		request := &chat.DouyinReceiveMessageRequest{
			ToUserId: userId,
		}
		response, _ := client.AcceptChatMessage(context.Background(), request)
		if len(response.ToUserIds) > 0 {
			fmt.Println("----- User ", userId, "Receive ", len(response.ToUserIds), " Messages -----")
			for index, _ := range response.ToUserIds {
				fmt.Println("Message", index+1, ":", response.FromUserIds[index], "->", response.ToUserIds[index], "——", response.Contents[index])
			}
		}
	}
}
