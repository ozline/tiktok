package test

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
)

var conn messageservice.Client
var token string

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := messageservice.NewClient("chat", client.WithHostPorts("43.136.122.18:8891"))
	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
