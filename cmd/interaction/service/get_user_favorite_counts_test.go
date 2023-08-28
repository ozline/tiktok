package service

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func TestGetUserFavoriteCounts(t *testing.T) {
	req := &interaction.UserFavoriteCountRequest{
		UserId: userId,
		Token:  token,
	}

	_, err := interactionService.GetUserFavoriteCount(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
