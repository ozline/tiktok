package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testUserFavoriteCount(t *testing.T) {
	req := &interaction.UserFavoriteCountRequest{
		UserId: userId,
		Token:  token,
	}

	_, err := interactionService.GetUserFavoriteCount(req)
	if err != nil {
		t.Logf("err [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkUserFavoriteCount(b *testing.B) {
	req := &interaction.UserFavoriteCountRequest{
		UserId: userId,
		Token:  token,
	}
	for n := 0; n < b.N; n++ {
		_, err := interactionService.GetUserFavoriteCount(req)
		if err != nil {
			b.Error(err)
		}
	}
}
