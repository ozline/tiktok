package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/follow"
)

func testFollowerCount(t *testing.T) {
	_, err := followService.FollowerCount(&follow.FollowerCountRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
