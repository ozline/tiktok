package main

import (
	"fmt"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
)

var conn videoservice.Client
var token string

func TestMain(m *testing.M) {

	config.Init("/home/jiuxia/tiktok/config", constants.VideoServiceName)
	fmt.Println(config.Server.Secret)
	// 连接服务器
	c, err := videoservice.NewClient("video", client.WithHostPorts("127.0.0.1:10006"))
	if err != nil {
		panic(err)
	}

	conn = c
	token, _ = utils.CreateToken(10000)
	m.Run()
}
