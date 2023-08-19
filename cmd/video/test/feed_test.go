package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
)

func TestFeed(t *testing.T) {
	req := &video.FeedRequest{
		//测试的时候根据数据库改
		LatestTime: 1692456150}
	resp, err := conn.Feed(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))

	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if resp.Base.Code != errno.SuccessCode {
		t.Error(errno.NewErrNo(resp.Base.Code, resp.Base.Msg))
		t.Fail()
	}
	if resp.VideoList == nil {
		t.Error(errno.ServiceError.WithMessage("videoList is null"))
		t.Fail()
	}

	fmt.Printf("Resp:\n%+v\n", resp)

}
