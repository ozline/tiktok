package main

import (
	"github.com/ozline/tiktok/pkg/constants"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var conn userservice.Client
var token string

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := userservice.NewClient("user",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:8888"))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
