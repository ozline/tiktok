package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func TestFavoriteList(t *testing.T) {

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
