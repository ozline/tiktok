package main

import (
	"testing"

	"github.com/ozline/tiktok/pkg/constants"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/follow/followservice"
)

var conn followservice.Client
var token string

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := followservice.NewClient("follow",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:8890"))

	if err != nil {
		panic(err)
	}

	conn = c
	m.Run()
}
