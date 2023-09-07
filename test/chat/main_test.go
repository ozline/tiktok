package test

import (
	"context"
	"testing"
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/cmd/chat/service"
	"github.com/ozline/tiktok/config"
)

var (
	from_user_id int64  = 3
	to_user_id   int64  = 2
	content_get  string = "test get"
	content_post string = "test post"
	create_at    string = time.Now().Format(time.RFC3339)
	chatService  *service.ChatService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()

	chatService = service.NewChatService(context.Background())
	m.Run()
}
func TestMainOrder(t *testing.T) {
	t.Run("rpc", testRpc)

	t.Run("get_message", testGetMessage)

	t.Run("post_message", testPostMessage)

	t.Run("get_message", testGetMessage)

	t.Run("polling", testPolling)

	t.Run("db", testDB)

	t.Run("redis", testRedis)
}

func BenchmarkMainOrder(b *testing.B) {
	b.Run("get_post_message", benchmarkGetAndPostMessage)

	b.Run("post_message", benchmarkPostMessage)
}
