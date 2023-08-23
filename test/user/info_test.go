package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/user"
)

func testGetUserInfo(t *testing.T) {
	_, err := userService.GetUser(&user.InfoRequest{
		UserId: id,
		Token:  token,
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
