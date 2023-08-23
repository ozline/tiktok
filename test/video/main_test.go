package main

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/video/dal"
	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/utils"
)

var (
	videoService *service.VideoService
	videoId      []int64
	token        string
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	cache.Init()
	videoService = service.NewVideoService(context.Background())
	token, _ = utils.CreateToken(10000)
	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("Feed", testFeed)
	t.Run("Get publish", testGetPublishVideo)
	t.Run("Get favorite", testGetFavoriteVideo)
	t.Run("RPC Test", testRPC)
}
