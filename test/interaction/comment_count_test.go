package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testCommentCount(t *testing.T) {
	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}
	_, err := interactionService.CountComments(req)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testCommentCount success---------------")
}

func benchmarkCommentCount(b *testing.B) {
	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}

	for i := 0; i < b.N; i++ {

		interactionService.CountComments(req)
	}
}
