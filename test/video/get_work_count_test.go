package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/video"
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
