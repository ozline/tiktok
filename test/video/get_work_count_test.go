package main

import (
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
)

func testWorkCount(t *testing.T) {
	_, err := videoService.GetWorkCount(&video.GetWorkCountRequest{
		UserId: 10000,
		Token:  token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
