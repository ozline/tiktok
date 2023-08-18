package main

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
)

func TestFavoriteAction(t *testing.T) {
	req := &interaction.FavoriteActionRequest{
		VideoId:    "1",
		ActionType: "1",
		Token:      token,
	}

	resp, err := conn.FavoriteAction(context.Background(), req, callopt.WithRPCTimeout(constants.RPCTimeout))
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
