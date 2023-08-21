package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
)

func testPostMessage(t *testing.T) {
	t.Log("------------testPostMessage Start---------------")
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		klog.Info(err)
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	req := &chat.MessagePostRequest{
		Token:      token,
		FromUserId: from_user_id,
		ToUserId:   to_user_id,
		Content:    content_post,
		CreateTime: &now,
	}
	err = chatService.SendMessage(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(time.Second * 4)
	t.Log("------------testPostMessage End---------------")
}

func benchmarkPostMessage(b *testing.B) {
	b.Log("------------benchmarkPostMessage Start---------------")
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	// b.N = 10
	for i := 0; i < b.N; i++ {
		for i := 0; i < 2000; i++ {
			now := time.Now().Format("2006-01-02 15:04:05")
			req := &chat.MessagePostRequest{
				Token:      token,
				FromUserId: from_user_id,
				ToUserId:   to_user_id,
				Content:    content_post + "-->" + strconv.FormatInt(int64(i), 10),
				CreateTime: &now,
			}
			err = chatService.SendMessage(req)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
		}
	}
	b.Log("------------benchmarkPostMessage End---------------")
}
