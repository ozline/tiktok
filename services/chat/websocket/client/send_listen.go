package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"

	//"golang.org/x/net/websocket"
	"log"
)

func main() {
	url := "ws://localhost:8888/ws" //服务器地址
	userInfo := make(map[string][]string)
	userInfo["UserName"] = append(userInfo["UserName"], "Send&Listen")
	ws, _, err := websocket.DefaultDialer.Dial(url, userInfo)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			err := ws.WriteMessage(websocket.BinaryMessage, []byte("ping"))
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
