package test

import (
	"strconv"
	"testing"
	"time"

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

func benchmarkGetAndPostMessage(b *testing.B) {
	token, _ := utils.CreateToken(2)
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: 3,
	}
	b.Log("---------------------------------")
	b.N = 1
	for i := 0; i < b.N; i++ {
		for i := 0; i < 2000; i++ {
			_, err := chatService.GetMessages(req, 2)
			if err != nil {
				b.Error(err)
				b.Fail()
			}
			if i%2000 == 0 {
				b.Log("----------------input-----------------")
			}
			now := time.Now().Unix()
			req_post := &chat.MessagePostRequest{
				Token:      token,
				FromUserId: 2,
				ToUserId:   3,
				Content:    "hello-->" + strconv.FormatInt(int64(i), 10),
				CreateTime: &now,
			}
			err = chatService.SendMessage(req_post)
			if err != nil {
				b.Error(err)
				b.Fail()
			}

		}
	}
	b.Log("------------TestGet success---------------")
}