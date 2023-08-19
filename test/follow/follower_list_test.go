package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/follow"
)

func testFollowerList(t *testing.T) {
	_, err := followService.FollowerList(&follow.FollowerListRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
