package main

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func testFeed(t *testing.T) {
	_, _, _, _, err := videoService.FeedVideo(&video.FeedRequest{
		LatestTime: 1692472274,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
