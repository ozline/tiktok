package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
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

	resp, err := interactionService.CreateComment(req, userId)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	commentId = resp.Id
	t.Logf("commentId: [%v] \n", commentId)

	_, err = interactionService.DeleteComment(req, userId)

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	t.Log("------------testCommentAction success---------------")
}

func benchmarkCommentAction(b *testing.B) {
	req := &interaction.CommentActionRequest{
		VideoId:     videoId,
		CommentText: &commentText,
		CommentId:   &commentId,
		Token:       token,
	}

	for i := 0; i < b.N; i++ {
		//interactionService.MatchSensitiveWords(commentText)

		resp, _ := interactionService.CreateComment(req, userId)

		commentId = resp.Id

		_, err := interactionService.DeleteComment(req, userId)
		if err != nil {
			b.Error(err)
		}

	}
}
