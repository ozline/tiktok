package main

import (
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"strconv"
	"testing"
)

func testCommentAction(t *testing.T) {
	_, err := interactionService.MatchSensitiveWords(commentText)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	req := &interaction.CommentActionRequest{
		VideoId:     videoId,
		CommentText: &commentText,
		CommentId:   &commentId,
		Token:       token,
	}

	resp, err := interactionService.CreateComment(req)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	commentId = strconv.FormatInt(resp.Id, 10)
	t.Logf("commentId: [%v] \n", commentId)

	_, err = interactionService.DeleteComment(req)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testCommentAction success---------------")
}
