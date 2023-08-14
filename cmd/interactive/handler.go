package main

import (
	"context"
	interactive "github.com/ozline/tiktok/kitex_gen/interactive"
)

// InteractiveServiceImpl implements the last service interface defined in the IDL.
type InteractiveServiceImpl struct{}

// FavoriteAction implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) FavoriteAction(ctx context.Context, req *interactive.FavoriteActionRequest) (resp *interactive.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) FavoriteList(ctx context.Context, req *interactive.FavoriteListRequest) (resp *interactive.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) CommentAction(ctx context.Context, req *interactive.CommentActionRequest) (resp *interactive.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the InteractiveServiceImpl interface.
func (s *InteractiveServiceImpl) CommentList(ctx context.Context, req *interactive.CommentListRequest) (resp *interactive.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
