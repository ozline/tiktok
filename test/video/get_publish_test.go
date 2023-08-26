package main

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func testGetPublishVideo(t *testing.T) {
	_, _, _, _, _, err := videoService.GetPublishVideoInfo(&video.GetPublishListRequest{
		UserId: 10000,
		Token:  token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
