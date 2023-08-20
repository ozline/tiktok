package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testFavoriteCount(t *testing.T) {

	req := &interaction.FavoriteCountRequest{
		VideoId: videoId,
		Token:   token,
	}

	_, err := interactionService.GetLikeCount(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testFavoriteCount success---------------")
}
