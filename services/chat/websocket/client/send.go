package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"strconv"
	"time"
	//"golang.org/x/net/websocket"
	"log"
)

func main() {
	url := "ws://localhost:8888/ws" //服务器地址
	userInfo := make(map[string][]string)
	userInfo["UserId"] = append(userInfo["UserId"], "1")
	ws, _, err := websocket.DefaultDialer.Dial(url, userInfo)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			message := WebMessage{
				FromUserId: 1,
				ToUserId:   2,
				Content:    "Hello Listen,This is Send",
				CreateTime: strconv.FormatInt(time.Now().UnixMicro(), 10), // 毫秒
			}
			str, _ := json.Marshal(message)
			err := ws.WriteMessage(websocket.BinaryMessage, str)

			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("receive: ", string(data))
	}
}
