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

func TestGetVideoIDByUid(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	videoService := service.NewVideoService(context.Background())
	_, err = videoService.GetVideoIDByUid(&video.GetVideoIDByUidRequset{Token: token, UserId: 10000})

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
