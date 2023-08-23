package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/follow"
)

func testFollowCount(t *testing.T) {
	_, err := followService.FollowCount(&follow.FollowCountRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
