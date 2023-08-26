package main

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func testFriendList(t *testing.T) {
	monkey.Patch(rpc.GetUser, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: touserid}, nil
	})

	monkey.Patch(rpc.GetMessage, func(ctx context.Context, req *chat.MessageListRequest, uid, tid int64) (string, int64, error) {
		return "hello", 1, nil
	})

	defer monkey.UnpatchAll()
	_, err := followService.FriendList(&follow.FriendListRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkFriendList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := followService.FriendList(&follow.FriendListRequest{
			UserId: id,
			Token:  token,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}

		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}
