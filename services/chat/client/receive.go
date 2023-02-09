package main

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gorilla/websocket"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat/tiktokchatservice"
	"strconv"
	"time"

	"context"
	"encoding/json"
	"log"
)

func main() {
	userId := 2
	ch := receiveMsgFromMySQL(int64(userId))
	time.Sleep(1 * time.Second)
	ch <- true
	close(ch)
	url := "ws://localhost:8888/ws" //服务器地址
	userInfo := make(map[string][]string)
	userInfo["UserId"] = append(userInfo["UserId"], "2")
	ws, _, err := websocket.DefaultDialer.Dial(url, userInfo)
	if err != nil {
		log.Fatal(err)
	}
	for {
		message := WebMessage{}
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(data), &message)
		//fmt.Println("receive: ", string(data))
		createtime, err := strconv.ParseInt(message.CreateTime, 10, 64)
		wasteTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(time.Now().UnixMicro()-createtime)/float64(1000)), 64) // 保留5位小数
		fmt.Println("Translate Time:", wasteTime, "ms")
	}
}

func receiveMsgFromMySQL(userId int64) chan bool {
	ticker := time.NewTicker(10 * time.Millisecond)
	stopChan := make(chan bool)
	go func(userId int64) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
				if err != nil {
					log.Fatal(err)
				}
				request := &chat.DouyinReceiveMessageRequest{
					ToUserId: userId,
				}
				response, _ := client.AcceptChatMessage(context.Background(), request)
				if response.StatusCode == 1 {
					fmt.Println("----- User ", userId, "Receive ", len(response.ToUserIds), " Messages -----")
					for index, _ := range response.ToUserIds {
						wasteTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(time.Now().UnixMicro()-response.CreateTime[index])/float64(1000)), 64) // 保留5位小
						fmt.Println("Message", index+1, ":", response.FromUserIds[index], "->", response.ToUserIds[index], "——", response.Contents[index], wasteTime, "ms")
					}
				}
			case stop := <-stopChan:
				if stop {
					return
				}
			}
		}
	}(userId)
	return stopChan
}
