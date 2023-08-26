package main

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func testCommentCount(t *testing.T) {
	monkey.Patch(rpc.UserInfo, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: userId}, nil
	})

	defer monkey.UnpatchAll()

	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}
	_, err := interactionService.CountComments(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testCommentCount success---------------")
}

func benchmarkCommentCount(b *testing.B) {
	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}

	for i := 0; i < b.N; i++ {
		_, err := interactionService.CountComments(req)
		if err != nil {
			b.Error(err)
		}
	}
}
