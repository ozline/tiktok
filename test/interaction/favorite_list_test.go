package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

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
