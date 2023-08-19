package main

import (
	"context"
	"github.com/ozline/tiktok/cmd/interaction/dal"
	"github.com/ozline/tiktok/cmd/interaction/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/utils"
	"testing"
)

var (
	videoId     string
	token       string
	commentText string
	commentId   string

	interactionService *service.InteractionService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()

	interactionService = service.NewInteractionService(context.Background())

	token, _ = utils.CreateToken(10000)
	commentText = "发条评论看看"
	videoId = "1"

	m.Run()
}

func TestMainOrder(t *testing.T) {

	t.Run("comment action", testCommentAction)

	t.Run("comment list", testCommentList)

	t.Run("comment count", testCommentCount)

	t.Run("RPC Test", testRPC)
}
