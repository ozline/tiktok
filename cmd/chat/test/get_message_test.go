package test

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func TestGetMessage(t *testing.T) {
	token, _ := utils.CreateToken(2)
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: 3,
	}

	resp, err := conn.MessageList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	klog.Info(resp)
}
