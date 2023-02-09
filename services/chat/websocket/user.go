package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

type User struct {
	addr     string
	userName string
	wsConn   *websocket.Conn
	sendChan chan []byte
}

func NewUser(clientInfo WebClient) {
	var user = &User{
		addr:     clientInfo.conn.RemoteAddr().String(),
		userName: clientInfo.UserName,
		wsConn:   clientInfo.conn,
		sendChan: make(chan []byte),
	}
	fmt.Println("----- Hello -----")
	user.online()
	go user.recvMessage()
	go user.sendMessage()
}

func (user *User) recvMessage() {
	defer user.offline()
	for {
		message := WebMessage{}
		_, p, err := user.wsConn.ReadMessage()
		json.Unmarshal(p, &message)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("From=", message.FromUser)
		fmt.Println("To=", message.ToUser)
		fmt.Println("Content=", message.Content)
		fmt.Println("CreateTime=", message.CreateTime)
		sendUser := server.onlineUserMap[message.ToUser]
		sendUser.sendChan <- p
		log.Printf("Recv [%s] msg:%s", user.addr, p)
	}
}

func (user *User) sendMessage() {
	defer user.offline()
	for {
		buf := <-user.sendChan

		err := user.wsConn.WriteMessage(1, buf)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Send [%s] msg:%s", user.addr, buf)
	}
}

func (user *User) online() {
	server.addOnlineUserMap(user)

	log.Printf("[%s] 上线了", user.addr)
}

func (user *User) offline() {
	user.wsConn.Close()
	server.deleteOnlineUserMap(user)

	log.Printf("[%s] 下线了", user.addr)
}
