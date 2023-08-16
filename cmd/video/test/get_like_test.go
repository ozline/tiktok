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
		VideoId: []int64{479434630982795264},
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
		t.Error(errno.NewErrNo(resp.Base.Code, resp.Base.Msg))
		t.Fail()
	}
	//0值和nil不会显示
	fmt.Printf("Resp:\n%+v\n", resp)
	fmt.Println("success")

}
