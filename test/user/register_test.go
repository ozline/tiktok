package main

import (
	"context"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
)

func testRegister(t *testing.T) {

	resp, err := conn.Register(context.Background(), &user.RegisterRequest{
		Username: username,
		Password: password,
	}, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	if resp.Base.Code != errno.SuccessCode {
		t.Error(errno.NewErrNo(resp.Base.Code, resp.Base.Msg))
	}

	// t.Logf("Resp:\n%v\n\n", resp)
}
