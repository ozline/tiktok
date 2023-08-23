package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testFavoriteList(t *testing.T) {

	req := &interaction.FavoriteListRequest{
		UserId: userId,
		Token:  token,
	}

	_, err := interactionService.FavoriteList(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testFavoriteList success---------------")
}

func benchmarkFavoriteList(b *testing.B) {
	req := &interaction.FavoriteListRequest{
		UserId: userId,
		Token:  token,
	}
	for n := 0; n < b.N; n++ {
		_, err := interactionService.FavoriteList(req)
		if err != nil {
			b.Error(err)
		}
	}
}
