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
	cli, err := messageservice.NewClient("chat",
		client.WithMuxConnection(constants.MuxConnection),
		client.WithHostPorts("0.0.0.0:10003"))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	resp, err := cli.MessageList(context.Background(), &chat.MessageListRequest{
		Token:    token,
		ToUserId: to_user_id,
	})
	if err != nil {
		t.Logf("result ==> %v", resp)
		t.Error(err)
		t.Fail()
	}
	t.Logf("result ==> %v", resp.MessageList)
	t.Logf("result ==> %v", resp.Base)
	t.Logf("result ==> %v", resp.Total)
	t.Logf("result ==> %v", len(resp.MessageList))
}
