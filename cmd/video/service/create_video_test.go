package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ozline/tiktok/cmd/video/dal"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/cmd/video/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

func TestCreateVideo(t *testing.T) {
	config.InitForTest()
	dal.Init()
	rpc.Init()
	videoService := service.NewVideoService(context.Background())
	// 测试token不存在导致的errno.AuthorizationFailedError错误
	_, err := videoService.CreateVideo(&video.PutVideoRequest{VideoFile: nil,
		Title: "test_title",
		Token: "",
	}, "test_video_URL", "test_cover_URL")
	if !errors.Is(err, errno.AuthorizationFailedError) {
		t.Error(err)
		t.Fail()
	}
	// 正常流程测试
	token, err := utils.CreateToken(10000)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	_, err = videoService.CreateVideo(&video.PutVideoRequest{VideoFile: nil,
		Title: "test_title",
		Token: token,
	}, "test_video_URL", "test_cover_URL")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
