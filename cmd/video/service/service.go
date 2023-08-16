package service

import "context"

type VideoService struct {
	ctx context.Context
}

// NewVideoService new VideoService
func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}
