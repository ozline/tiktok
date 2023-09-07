package main

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/interaction/dal"
	"github.com/ozline/tiktok/cmd/interaction/service"
	"github.com/ozline/tiktok/config"
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
	dal.Init("../../config")
	interactionService = service.NewInteractionService(context.Background())

	userId = 10000
	token, _ = utils.CreateToken(userId)
	commentText = "发条评论看看"

	videoId = 1
	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("comment list", testCommentList)

	t.Run("comment count", testCommentCount)

	t.Run("favorite count", testVideoFavoriteCount)

	t.Run("user favorite count", testUserFavoriteCount)

	t.Run("is favorite", testIsFavorite)

	t.Run("RPC Test", testRPC)
}

func BenchmarkMainOrder(b *testing.B) {
	b.Run("comment action", benchmarkCommentAction)

	b.Run("comment list", benchmarkCommentList)

	b.Run("comment count", benchmarkCommentCount)

	b.Run("favorite action", benchmarkFavoriteAction)

	b.Run("favorite list", benchmarkFavoriteList)

	b.Run("favorite count", benchmarkFavoriteVideoCount)

	b.Run("user favorite count", benchmarkUserFavoriteCount)

	b.Run("user total favorited", benchmarkUserTotalFavorited)

	b.Run("is favorite", benchmarkIsFavorite)
}
