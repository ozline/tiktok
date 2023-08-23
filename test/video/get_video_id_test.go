package main

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func testGetVideoIDByUid(t *testing.T) {
	_, err := videoService.GetVideoIDByUid(&video.GetVideoIDByUidRequset{
		UserId: 10000,
		Token:  token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
