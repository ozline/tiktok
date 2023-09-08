package test

import (
	"fmt"
	"strconv"
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
	req := &chat.MessagePostRequest{
		Token:    token,
		ToUserId: to_user_id,
		Content:  fmt.Sprintf("test post %d", time.Now().Unix()),
	}
	err = chatService.SendMessage(req, from_user_id, create_at)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	time.Sleep(time.Second * 4)
}

func benchmarkPostMessage(b *testing.B) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	// b.N = 10
	for i := 0; i < b.N; i++ {
		for i := 0; i < 2000; i++ {
			req := &chat.MessagePostRequest{
				Token:    token,
				ToUserId: to_user_id,
				Content:  content_post + "-->" + strconv.FormatInt(int64(i), 10),
			}
			err = chatService.SendMessage(req, from_user_id, create_at)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
		}
	}
}
