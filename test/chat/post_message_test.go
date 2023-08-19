package test

import (
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testPostMessage(t *testing.T) {
	token, err := utils.CreateToken(2)
	if err != nil {
		klog.Info(err)
	}
	now := time.Now().Unix()
	req := &chat.MessagePostRequest{
		Token:      token,
		FromUserId: 2,
		ToUserId:   3,
		Content:    "hellooooooooooooooooooooo",
		CreateTime: &now,
	}
	t.Log("-------------------------------------------")
	err = chatService.SendMessage(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(time.Second * 4)
	t.Log("------------TestPostMessage success---------------")
}
