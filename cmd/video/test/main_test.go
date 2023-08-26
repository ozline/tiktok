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
	c, err := videoservice.NewClient("video", client.WithHostPorts("127.0.0.1:10006"))

	if err != nil {
		panic(err)
	}
	token, _ = utils.CreateToken(10000)
	conn = c
	m.Run()
}
