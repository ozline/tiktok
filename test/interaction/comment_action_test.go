package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

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
