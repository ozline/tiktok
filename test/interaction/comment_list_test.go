package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func testCommentList(t *testing.T) {
	req := &interaction.CommentListRequest{
		VideoId: videoId,
		Token:   token,
	}

	_, err := interactionService.GetComments(req)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testCommentList success---------------")
}

func benchmarkCommentList(b *testing.B) {
	req := &interaction.CommentListRequest{
		VideoId: videoId,
		Token:   token,
	}

	for i := 0; i < b.N; i++ {

		_, err := interactionService.GetComments(req)
		if err != nil {
			b.Error(err)
		}
	}
}
