package test

import (
	"strconv"
	"testing"

	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/utils"
	"golang.org/x/sync/errgroup"
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
	t.Log(resp)
	for _, v := range resp {
		t.Log(v)
	}
}
func testPolling(t *testing.T) {
	token, err := utils.CreateToken(from_user_id)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	req := &chat.MessageListRequest{
		Token:    token,
		ToUserId: to_user_id,
	}
	res := make([]*db.Message, 0)
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			resp, err := chatService.GetMessages(req, from_user_id)
			if err != nil {
				return err
			}
			res = append(res, resp...)
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		t.Error(err)
		t.Fail()
	}
	for _, v := range res {
		t.Log("=======================")
		t.Log(v.Content)
	}
	t.Log("count===>", len(res))
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
