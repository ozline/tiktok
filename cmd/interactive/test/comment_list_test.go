package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"github.com/ozline/tiktok/pkg/errno"
	"testing"
	"time"
)

func TestCommentList(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMDAwMCwiZXhwIjoxNjkyNjE3NzQ5LCJpYXQiOjE2OTIwMTI5NDksImlzcyI6InRpa3RvayJ9.NS6ESu4WVohWJ7lGMBn6PiiKrnnsB67PHq8EtcYw5n8"
	req := &interactive.CommentListRequest{
		VideoId: "1",
		Token:   token,
	}

	resp, err := conn.CommentList(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

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
