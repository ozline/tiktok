package test

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/chat/messageservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
)

func testRpc(t *testing.T) {
	_, err := messageservice.NewClient("chat",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10003"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
