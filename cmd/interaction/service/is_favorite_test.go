package service

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func TestIsFavorite(t *testing.T) {
	req := &interaction.IsFavoriteRequest{
		UserId:  userId,
		VideoId: videoId,
		Token:   token,
	}

	_, err := interactionService.IsFavorite(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
