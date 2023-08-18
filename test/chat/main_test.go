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
	t.Run("getMessage", testGetMessage)

	t.Run("postMessage", testPostMessage)

	t.Run("db", testDB)

	t.Run("redis", testRedis)
}
