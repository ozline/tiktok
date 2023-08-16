package main

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/errno"
)

func TestFriendList(t *testing.T) {
	req := &follow.FriendListRequest{
		UserId: 10001,
		Token:  token,
	}

	resp, err := conn.FriendList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if resp.Base.Code != errno.SuccessCode {
		t.Error(errno.NewErrNo(resp.Base.Code, *resp.Base.Msg))
		t.Fail()
	}

	t.Logf("Resp:\n%v\n\n", resp)
}
