package main

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/ozline/tiktok/cmd/interaction/dal"
	"github.com/ozline/tiktok/cmd/interaction/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
)

var (
	userId      int64
	videoId     int64
	token       string
	commentText string
	commentId   int64

	interactionService *service.InteractionService
)

func TestMain(m *testing.M) {

	config.InitForTest()
	dal.Init()

	interactionService = service.NewInteractionService(context.Background())

	token, _ = utils.CreateToken(10000)
	commentText = "发条评论看看"
	userId = 1
	videoId = 1
	m.Run()
}

func TestMainOrder(t *testing.T) {

	t.Run("comment action", testCommentAction)

	t.Run("comment list", testCommentList)

	t.Run("comment count", testCommentCount)

	t.Run("favorite action", TestFavoriteAction)

	t.Run("favorite list", TestFavoriteList)

	t.Run("favorite count", TestFavoriteCount)

	t.Run("RPC Test", testRPC)
}
