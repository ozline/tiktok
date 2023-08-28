package service

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func TestDislike(t *testing.T) {
	req := &interaction.FavoriteActionRequest{
		VideoId:    videoId,
		Token:      token,
		ActionType: 2,
	}

	err := interactionService.Dislike(req, userId)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
