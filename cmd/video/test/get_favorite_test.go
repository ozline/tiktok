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

func TestGetLikedVideo(t *testing.T) {
	req := &video.GetFavoriteVideoInfoRequest{
		//测试的时候根据数据库改
		VideoId: []int64{480401105549787136},
		Token:   token,
	}
	resp, err := conn.GetFavoriteVideoInfo(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
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
