package test

import (
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testGetMessage(t *testing.T) {
	token, _ := utils.CreateToken(2)
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: 3,
	}

	resp, err := chatService.GetMessages(req, 2)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	for _, v := range resp {
		klog.Info("-----------------------------")
		klog.Info(v)
	}
	t.Log("------------TestGet success---------------")

}
