package main

import (
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"testing"
)

func testCommentCount(t *testing.T) {

	req := &interaction.CommentCountRequest{
		VideoId: "1",
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
