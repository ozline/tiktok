package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
)

var conn videoservice.Client
var token string

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := videoservice.NewClient("video", client.WithHostPorts("0.0.0.0:8892"))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
