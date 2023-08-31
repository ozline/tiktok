package service

import (
	"context"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/config"
)

var (
	chatservice *ChatService
	create_at   string = time.Now().Format(time.RFC3339)
	ac_type     int64  = 1
)

const (
	content      string = "cover test"
	from_user_id int64  = 2
	to_user_id   int64  = 3
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	chatservice = NewChatService(context.Background())
	m.Run()
}
func TestMainOrder(t *testing.T) {
	t.Run("post_message", testPostMessage)

	t.Run("get_message", testGetMessage)
}
