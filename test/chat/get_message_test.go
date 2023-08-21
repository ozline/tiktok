package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testGetMessage(t *testing.T) {
	t.Log("------------testGetMessage Start---------------")
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
	for _, v := range resp {
		t.Log(("-----------------------------"))
		t.Log(v)
	}
	t.Log("------------testGetMessage End---------------")

}

func benchmarkGetAndPostMessage(b *testing.B) {
	b.Log("-----benchmarkGetAndPostMessage Start-----")
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: to_user_id,
	}

	// b.N = 1
	for i := 0; i < b.N; i++ {
		b.Logf("----epoch:%v ----", i)
		for i := 0; i < 2000; i++ {
			_, err := chatService.GetMessages(req, from_user_id)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			req_post := &chat.MessagePostRequest{
				Token:      token,
				FromUserId: from_user_id,
				ToUserId:   to_user_id,
				Content:    content_get + "-->" + strconv.FormatInt(int64(i), 10),
				CreateTime: &now,
			}
			err = chatService.SendMessage(req_post)
			if err != nil {
				b.Error(err)
				b.Fail()
			}

		}
	}
	b.Log("-----benchmarkGetAndPostMessage End-----")
}
