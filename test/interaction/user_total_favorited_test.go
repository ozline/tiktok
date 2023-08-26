package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

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
