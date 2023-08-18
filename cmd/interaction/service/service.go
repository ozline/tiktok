package service

import "context"

type InteractionService struct {
	ctx context.Context
}

// NewinteractionService new interactionService
func NewInteractionService(ctx context.Context) *InteractionService {
	return &InteractionService{ctx: ctx}
}
