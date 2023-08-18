package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
)

var conn interactionservice.Client
var token string

func TestMain(m *testing.M) {
	token, _ = utils.CreateToken(10000)
	// 连接服务器
	c, err := interactionservice.NewClient("interaction",
		client.WithHostPorts("0.0.0.0:8889"),
		client.WithMuxConnection(constants.MuxConnection))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
