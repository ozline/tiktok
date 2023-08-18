package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"github.com/ozline/tiktok/pkg/errno"
	"testing"
	"time"
)

func TestCommentAction(t *testing.T) {
	content := "呜呜！！！"
	id := "480406139545059328"
	req := &interactive.CommentActionRequest{
		VideoId:     "1",
		ActionType:  "2",
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
