package main

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/follow/dal"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/cmd/follow/service"
	"github.com/ozline/tiktok/config"
)

var (
	touserid   int64
	actiontype int64
	token      string
	id         int64

	followService *service.FollowService
)

func TestMain(m *testing.M) {
	config.InitForTest()
	dal.Init()
	rpc.Init()

	followService = service.NewFollowService(context.Background())

	touserid = 10002
	actiontype = 1
	id = 10001

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("action", testAction)

	t.Run("followList", testFollowList)

	t.Run("followerList", testFollowerList)

	t.Run("friendList", testFriendList)

	t.Run("followCount", testFollowCount)

	t.Run("followerCount", testFollowerCount)

	t.Run("isfollow", testIsFollow)

	t.Run("RPC Test", testRPC)
}

func BenchmarkMainOrder(b *testing.B) {
	b.Run("action", benchmarkAction)

	b.Run("followList", benchmarkFollowList)

	b.Run("followerList", benchmarkFollowerList)

	b.Run("friendList", benchmarkFriendList)
}
