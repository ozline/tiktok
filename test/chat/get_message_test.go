package test

import (
	"strconv"
	"testing"

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
	resp, err := chatService.GetMessages(req, from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("---------------------------")
	t.Log(resp)
	for _, v := range resp {
		t.Log("---------------------------")
		t.Log(v)
		t.Log("---------------------------")
	}
}

func benchmarkGetAndPostMessage(b *testing.B) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: to_user_id,
	}
	for i := 0; i < b.N; i++ {
		b.Logf("----epoch:%v ----", i)
		for i := 0; i < 100; i++ {
			_, err := chatService.GetMessages(req, from_user_id)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
			req_post := &chat.MessagePostRequest{
				Token:    token,
				ToUserId: to_user_id,
				Content:  content_get + "-->" + strconv.FormatInt(int64(i), 10),
			}
			err = chatService.SendMessage(req_post, from_user_id, create_at)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
		}
	}
}
