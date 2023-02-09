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
	userInfo["UserName"] = append(userInfo["UserName"], "Listen")

	ws, _, err := websocket.DefaultDialer.Dial(url, userInfo)
	if err != nil {
		log.Fatal(err)
	}

	for {

		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		message := WebMessage{}
		err = json.Unmarshal(data, &message)
		if err != nil {
			fmt.Println("反序列化失败")
		}
		//fmt.Println("receive: ", string(data))
		wasteTime, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(time.Now().UnixMicro()-message.CreateTime)/float64(1000)), 64) // 保留5位小
		fmt.Println("Translate Time=", wasteTime, "ms")
	}
}
