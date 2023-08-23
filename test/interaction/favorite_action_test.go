package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testFavoriteAction(t *testing.T) {
	req := &interaction.FavoriteActionRequest{
		VideoId: videoId,
		Token:   token,
	}

	err := interactionService.Like(req, userId)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	err = interactionService.Dislike(req, userId)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testFavoriteAction success---------------")
}

func benchmarkFavoriteAction(b *testing.B) {
	req := &interaction.FavoriteActionRequest{
		VideoId: videoId,
		Token:   token,
	}
	for n := 0; n < b.N; n++ {
		err := interactionService.Like(req, userId)
		if err != nil {
			b.Error(err)
		}
		err = interactionService.Dislike(req, userId)
		if err != nil {
			b.Error(err)
		}
	}
}
