package main

import (
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 可利用wsConn读写数据与客户端通信
	userInfo := make(map[string][]string)
	wsConn, err := upgrader.Upgrade(w, r, userInfo) //将 http 升级到 WebSocket 协议
	//fmt.Println("--- Server ---", r.Header.Get("UserName"))
	userid, _ := strconv.ParseInt(r.Header.Get("UserId"), 10, 64)
	seqid, _ := strconv.ParseInt(r.Header.Get("SeqId"), 10, 64)
	clientInfo := WebClient{
		UserId: userid,
		conn:   wsConn,
		seqID:  seqid,
	}
	if err != nil {
		log.Println(err)
	}

	NewUser(clientInfo)
}

func NewUser(clientInfo WebClient) {
	var user = &User{
		addr:     clientInfo.conn.RemoteAddr().String(),
		userId:   clientInfo.UserId,
		wsConn:   clientInfo.conn,
		sendChan: make(chan []byte),
		ackID:    1,
	}
	//fmt.Println("----- New Client,UserId=", user.userId, ",UserAddr=", user.addr, " -----")
	user.online()
	go user.recvMessage()
	go user.sendMessage()
}

func (server *Server) addOnlineUserMap(user *User) {
	server.userMapLock.Lock()
	if _, ok := server.onlineUserMap[user.userId]; ok {
		server.onlineUserMap[user.userId].wsConn = user.wsConn
	} else {
		server.onlineUserMap[user.userId] = user
	}

	server.userMapLock.Unlock()
}

func (server *Server) deleteOnlineUserMap(user *User) {
	server.userMapLock.Lock()
	delete(server.onlineUserMap, user.userId)
	server.userMapLock.Unlock()
}
