package main

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"testing"
)

func TestCommentCount(t *testing.T) {

	req := &interaction.CommentCountRequest{
		VideoId: "1",
		Token:   &token,
	}

	resp, err := conn.CommentCount(context.Background(), req, callopt.WithRPCTimeout(constants.RPCTimeout))
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
