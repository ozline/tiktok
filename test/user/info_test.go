package main

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func testGetUserInfo(t *testing.T) {

	token, err := utils.CreateToken(10001)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	req := &user.InfoRequest{
		UserId: id, // 按需修改账号
		Token:  token,
	}

	resp, err := conn.Info(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if resp.Base.Code != errno.SuccessCode {
		t.Error(errno.NewErrNo(resp.Base.Code, resp.Base.Msg))
		t.Fail()
	}
}
