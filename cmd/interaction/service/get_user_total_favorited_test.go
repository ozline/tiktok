package service

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func TestGetUserTotalFavorited(t *testing.T) {
	videoIDList := make([]int64, 0)
	videoIDList = append(videoIDList, videoId)
	monkey.Patch(rpc.GetUserVideoList, func(ctx context.Context, req *video.GetVideoIDByUidRequset) ([]int64, error) {
		return videoIDList, nil
	})
	defer monkey.UnpatchAll()

	req := &interaction.UserTotalFavoritedRequest{
		Token:  token,
		UserId: userId,
	}

	_, err := interactionService.GetUserTotalFavorited(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
