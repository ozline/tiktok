package main

import (
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/user/userservice"
	"github.com/ozline/tiktok/pkg/constants"
)

func testRPC(t *testing.T) {
	_, err := userservice.NewClient("video",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10006"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
