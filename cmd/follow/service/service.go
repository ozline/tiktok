package service

import "context"

type FollowService struct {
	ctx context.Context
}

// NewFollowService new FollowService
func NewFollowService(ctx context.Context) *FollowService {
	return &FollowService{ctx: ctx}
}
