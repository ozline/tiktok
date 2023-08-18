package main

import (
	"testing"

	"github.com/ozline/tiktok/pkg/constants"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
)

var (
	conn     userservice.Client
	username string
	password string
	token    string
	id       int64
)

func TestMain(m *testing.M) {
	// 连接服务器
	c, err := userservice.NewClient("user",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10002"))

	if err != nil {
		panic(err)
	}

	username = "ozline"
	password = "123456"

	conn = c
	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("register", testRegister)

	t.Run("login", testLogin)

	t.Run("info", testGetUserInfo)

	t.Run("RPC Test", testRPC)
}
