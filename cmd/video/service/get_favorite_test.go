package service_test

import (
	"context"
	"testing"

	"github.com/ozline/tiktok/cmd/video/dal"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/cmd/video/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/utils"
)

func TestGetFavoriteVideoInfo(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	videoService := service.NewVideoService(context.Background())
	_, _, _, _, err = videoService.GetFavoriteVideoInfo(&video.GetFavoriteVideoInfoRequest{
		VideoId: []int64{482581113097682944, 483299894140862464, 483302572409487360},
		Token:   token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
