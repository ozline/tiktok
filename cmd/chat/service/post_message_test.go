package service

import (
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testPostMessage(t *testing.T) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		klog.Info(err)
	}
	now := time.Now().Format(time.DateTime)
	req := &chat.MessagePostRequest{
		Token:      token,
		FromUserId: from_user_id,
		ToUserId:   to_user_id,
		Content:    content,
		CreateTime: &now,
	}
	err = chatservice.SendMessage(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req.FromUserId = to_user_id
	req.ToUserId = from_user_id
	err = chatservice.SendMessage(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(2 * time.Second)
}
