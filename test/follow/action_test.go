package main

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/follow/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/utils"
)

func testAction(t *testing.T) {
	monkey.Patch(rpc.GetUser, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: touserid}, nil
	})

	defer monkey.UnpatchAll()

	token, _ = utils.CreateToken(id)
	err := followService.Action(&follow.ActionRequest{
		Token:      token,
		ToUserId:   touserid,
		ActionType: actiontype,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkAction(b *testing.B) {
	token, _ = utils.CreateToken(id)
	for i := 0; i < b.N; i++ {
		err := followService.Action(&follow.ActionRequest{
			Token:      token,
			ToUserId:   touserid,
			ActionType: actiontype,
		})

		if err != nil {
			b.Errorf("err: [%v] \n", err)
		}

		time.Sleep(100 * time.Millisecond) // Add a sleep to simulate some processing time
	}
}
