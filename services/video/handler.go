package main

import (
	"context"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct {
	DatabaseTable string
}

// PutVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	//videoInfo := req.VideoInfo

	return
}

// DeleteVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) DeleteVideo(ctx context.Context, req *video.DeleteVideoRequest) (resp *video.DeleteVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) GetVideo(ctx context.Context, req *video.GetVideoRequest) (resp *video.GetVideoResponse, err error) {
	// TODO: Your code here...
	return
}
