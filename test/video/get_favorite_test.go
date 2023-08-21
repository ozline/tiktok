package main

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func testGetLikedVideo(t *testing.T) {
	_, _, err := videoService.GetFavoriteVideoInfo(&video.GetFavoriteVideoInfoRequest{
		VideoId: videoId,
		Token:   token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
