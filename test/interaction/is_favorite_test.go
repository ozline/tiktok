package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testIsFavorite(t *testing.T) {
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

func benchmarkIsFavorite(b *testing.B) {
	req := &interaction.IsFavoriteRequest{
		UserId:  userId,
		VideoId: videoId,
		Token:   token,
	}
	for n := 0; n < b.N; n++ {
		_, err := interactionService.IsFavorite(req)
		if err != nil {
			b.Error(err)
		}
	}
}
