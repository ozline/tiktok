package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	onlineUserMap map[int64]*User
	userMapLock   sync.RWMutex
	sqlLayer      *websocket.Conn
}
type WebClient struct {
	UserId int64
	conn   *websocket.Conn
}

var server = Server{
	onlineUserMap: make(map[int64]*User),
}

var upgrader = websocket.Upgrader{ // 用于将http请求升级i为长连接的websocket
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 可利用wsConn读写数据与客户端通信
	userInfo := make(map[string][]string)
	wsConn, err := upgrader.Upgrade(w, r, userInfo) //将 http 升级到 WebSocket 协议
	//fmt.Println("--- Server ---", r.Header.Get("UserName"))
	userid, _ := strconv.ParseInt(r.Header.Get("UserId"), 10, 64)
	clientInfo := WebClient{
		UserId: userid,
		conn:   wsConn,
	}
	if err != nil {
		log.Println(err)
	}

	NewUser(clientInfo)
}

func (server *Server) Start() {
	http.HandleFunc("/", handler)     // 自定义处理函数
	http.ListenAndServe(":8888", nil) // 创建监听端口
}

func (server *Server) addOnlineUserMap(user *User) {
	server.userMapLock.Lock()
	server.onlineUserMap[user.userId] = user
	server.userMapLock.Unlock()
}

func (server *Server) deleteOnlineUserMap(user *User) {
	server.userMapLock.Lock()
	delete(server.onlineUserMap, user.userId)
	server.userMapLock.Unlock()
}
