package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testUserTotalFavorited(t *testing.T) {
	req := &interaction.UserTotalFavoritedRequest{
		UserId: userId,
		Token:  token,
	}

	_, err := interactionService.GetUserTotalFavorited(req)
	if err != nil {
		t.Logf("err [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}

func benchmarkUserTotalFavorited(b *testing.B) {
	req := &interaction.UserTotalFavoritedRequest{
		UserId: userId,
		Token:  token,
	}
	for n := 0; n < b.N; n++ {
		_, err := interactionService.GetUserTotalFavorited(req)
		if err != nil {
			b.Error(err)
		}
	}
}
