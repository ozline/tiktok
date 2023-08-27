package main

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func testCommentAction(t *testing.T) {
	monkey.Patch(rpc.UserInfo, func(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
		return &user.User{Id: userId}, nil
	})

	defer monkey.UnpatchAll()
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
		return
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
		// interactionService.MatchSensitiveWords(commentText)
		resp, _ := interactionService.CreateComment(req, userId)

		if resp == nil {
			continue
		}
		commentId = resp.Id

		_, err := interactionService.DeleteComment(req, userId)
		if err != nil {
			b.Error(err)
		}
	}
}
