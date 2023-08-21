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
	token, err := utils.CreateToken(3)
	if err != nil {
		klog.Info(err)
	}
	now := time.Now().Unix()
	req := &chat.MessagePostRequest{
		Token:      token,
		FromUserId: 3,
		ToUserId:   2,
		Content:    "test post",
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

func benchmarkPostMessage(b *testing.B) {
	token, err := utils.CreateToken(3)
	if err != nil {
		b.Log("err------->", err)
	}
	b.N = 10
	for i := 0; i < b.N; i++ {

		for i := 0; i < 2000; i++ {

			now := time.Now().Unix()
			req := &chat.MessagePostRequest{
				Token:      token,
				FromUserId: 3,
				ToUserId:   2,
				Content:    "hello-->" + strconv.FormatInt(int64(i), 10),
				CreateTime: &now,
			}
			err = chatService.SendMessage(req)
			if err != nil {
				b.Log("err------->", err)
			}
		}
	}

}
