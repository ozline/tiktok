package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func TestFavoriteList(t *testing.T) {
	videoIDList := make([]*video.Video, 0)
	videoIDList = append(videoIDList, &video.Video{
		Id: videoId,
	})
	monkey.Patch(rpc.GetFavoriteVideoList, func(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) ([]*video.Video, error) {
		return videoIDList, nil
	})
	defer monkey.UnpatchAll()

	req := &interaction.FavoriteListRequest{
		UserId: userId,
		Token:  token,
	}

	_, err := interactionService.FavoriteList(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
