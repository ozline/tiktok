package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/gorilla/websocket"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat/tiktokchatservice"
	"log"
	"strconv"
)

type User struct {
	addr     string
	userId   int64
	wsConn   *websocket.Conn
	sendChan chan []byte
}

func NewUser(clientInfo WebClient) {
	var user = &User{
		addr:     clientInfo.conn.RemoteAddr().String(),
		userId:   clientInfo.UserId,
		wsConn:   clientInfo.conn,
		sendChan: make(chan []byte),
	}
	fmt.Println("----- New Client,UserId=", user.userId, ",UserAddr=", user.addr, " -----")
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
		_, ok := server.onlineUserMap[message.ToUserId]
		//fmt.Println("ToUserId=", message.ToUserId, ",Map State=", ok)
		if _, ok = server.onlineUserMap[message.ToUserId]; ok {
			// 存在

			sendUser := server.onlineUserMap[message.ToUserId]
			sendUser.sendChan <- p
			fmt.Println("----- User ", message.ToUserId, " is Online -----")
		} else {
			fmt.Println("----- We Need to Connect to DataBase -----")
			user.sendMsgToSqlLayer(message)

		}
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

func (user *User) sendMsgToSqlLayer(message WebMessage) {
	client, err := tiktokchatservice.NewClient("kitexprotobuf", client.WithHostPorts("0.0.0.0:8891"))
	if err != nil {
		log.Fatal(err)
	}
	createtime, err := strconv.ParseInt(message.CreateTime, 10, 64)
	request := &chat.DouyinSendMessageRequest{
		FromUserId: message.FromUserId,
		ToUserId:   message.ToUserId,
		Content:    message.Content,
		CreateTime: createtime,
	}
	_, err = client.SendChatMessage(context.Background(), request)
	if err != nil {
		log.Fatal("error", err.Error())
	}

}
