package main

import (
	"testing"

	"github.com/ozline/tiktok/pkg/constants"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var conn userservice.Client
var token string

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := userservice.NewClient("user",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("127.0.0.1:10002"))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
