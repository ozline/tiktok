package main

import (
	"fmt"
	"testing"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/pack"
)

func testFeed(t *testing.T) {
	videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err := videoService.FeedVideo(&video.FeedRequest{
		LatestTime: 1693010318,
		Token:      token,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(pack.VideoList(videoList, userList, favoriteCountList, commentCountList, isFavoriteList))

}
