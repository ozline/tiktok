package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testVideoFavoriteCount(t *testing.T) {
	req := &interaction.VideoFavoritedCountRequest{
		VideoId: videoId,
		Token:   token,
	}

	_, err := interactionService.GetVideoFavoritedCount(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testFavoriteCount success---------------")
}

func benchmarkFavoriteVideoCount(b *testing.B) {
	req := &interaction.VideoFavoritedCountRequest{
		VideoId: videoId,
		Token:   token,
	}
	for n := 0; n < b.N; n++ {
		_, err := interactionService.GetVideoFavoritedCount(req)
		if err != nil {
			b.Error(err)
		}
	}
}
