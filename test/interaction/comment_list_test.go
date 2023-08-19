package main

import (
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"testing"
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
