package main

import (
	"testing"

	"github.com/ozline/tiktok/kitex_gen/interaction"
)

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
