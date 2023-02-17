package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var server = Server{
	onlineUserMap: make(map[int64]*User),
}
var upgrader = websocket.Upgrader{ // 用于将http请求升级i为长连接的websocket
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Server struct {
	onlineUserMap map[int64]*User
	userMapLock   sync.RWMutex
	sqlLayer      *websocket.Conn
}

type WebClient struct {
	UserId int64
	conn   *websocket.Conn
	seqID  int64
}

type User struct {
	addr     string
	userId   int64
	wsConn   *websocket.Conn
	sendChan chan []byte
	ackID    int64
}

func (server *Server) Start() {
	http.HandleFunc("/", handler)     // 自定义处理函数
	http.ListenAndServe(":8888", nil) // 创建监听端口
}
