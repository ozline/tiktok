package main

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
)

func TestCommentAction(t *testing.T) {
	content := "test！！！"
	id := "10001"
	req := &interaction.CommentActionRequest{
		VideoId:     "1",
		ActionType:  "1",
		CommentText: &content,
		CommentId:   &id,
		Token:       token,
	}

	resp, err := conn.CommentAction(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

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
