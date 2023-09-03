package service

import (
	"testing"
	"time"

	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testGetMessage(t *testing.T) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: to_user_id,
	}

	_, err = chatservice.GetMessages(req, from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	token, err = utils.CreateToken(to_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(1 * time.Second)
	req.Token = token
	req.ToUserId = from_user_id
	msg, err := chatservice.GetMessages(req, to_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	for _, v := range msg {
		t.Log(v)
	}
	time.Sleep(1 * time.Second)
}
