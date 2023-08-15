package main

import (
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/interactive/interactiveservice"
	"github.com/ozline/tiktok/pkg/constants"
	"testing"
)

var conn interactiveservice.Client

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := interactiveservice.NewClient("interactive",
		client.WithHostPorts("0.0.0.0:8889"),
		client.WithMuxConnection(constants.MuxConnection))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
