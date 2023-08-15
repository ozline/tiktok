package service

import "context"

type CommentService struct {
	ctx context.Context
}

// NewCommentService new CommentService
func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}
