package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/pkg/utils"
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
	token, _ = utils.CreateToken(10000)
	m.Run()
}
