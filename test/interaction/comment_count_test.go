package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testCommentCount(t *testing.T) {
	req := &interaction.CommentCountRequest{
		VideoId: videoId,
		Token:   &token,
	}
	count, err := interactionService.CountComments(req, 0)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Logf("count: [%v] \n", count)
	t.Log("------------testCommentCount success---------------")
}

func benchmarkCommentCount(b *testing.B) {
	req := &interaction.CommentCountRequest{
		VideoId: 1,
		Token:   &token,
	}

	for i := 0; i < b.N; i++ {
		_, err := interactionService.CountComments(req, 0)
		if err != nil {
			b.Error(err)
		}
	}
}
