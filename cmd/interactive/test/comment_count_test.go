package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"github.com/ozline/tiktok/pkg/errno"
	"testing"
	"time"
)

func TestCommentCount(t *testing.T) {
	req := &interactive.CommentCountRequest{
		VideoId: "1",
	}

	resp, err := conn.CommentCount(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if resp.Base.Code != errno.SuccessCode {
		t.Error(errno.NewErrNo(resp.Base.Code, *resp.Base.Msg))
		t.Fail()
	}

	t.Logf("Resp:\n%v\n\n", resp)
}
