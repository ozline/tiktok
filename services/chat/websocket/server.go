package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	onlineUserMap map[string]*User
	userMapLock   sync.RWMutex
}
type WebClient struct {
	UserName string
	conn     *websocket.Conn
}

var server = Server{
	onlineUserMap: make(map[string]*User),
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
	clientInfo := WebClient{
		UserName: r.Header.Get("UserName"),
		conn:     wsConn,
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
	server.onlineUserMap[user.userName] = user
	server.userMapLock.Unlock()
}

func (server *Server) deleteOnlineUserMap(user *User) {
	server.userMapLock.Lock()
	delete(server.onlineUserMap, user.addr)
	server.userMapLock.Unlock()
}
