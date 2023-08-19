package main

import (
	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
	"testing"
)

var conn interactionservice.Client
var token string

func TestMain(m *testing.M) {
	token, _ = utils.CreateToken(10000)
	// 连接服务器
	c, err := interactionservice.NewClient("interactive",
		client.WithHostPorts("0.0.0.0:10005"),
		client.WithMuxConnection(constants.MuxConnection))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
