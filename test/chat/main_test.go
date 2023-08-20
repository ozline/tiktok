package test

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/chat/dal"
	"github.com/ozline/tiktok/cmd/chat/service"
	"github.com/ozline/tiktok/config"
)

var (
	username string
	password string
	token    string
	id       int64

	chatService *service.ChatService
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

	t.Run("db", testDB)

	t.Run("redis", testRedis)

}

func BenchmarkMainOrder(b *testing.B) {

	b.Run("get_post_message", benchmarkGetAndPostMessage)

	b.Run("post_message", benchmarkPostMessage)
}
