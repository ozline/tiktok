package main

import (
	"context"
	"crypto/md5"
	video "github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"time"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct {
	DatabaseTable string
}

// PutVideo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	video := req.VideoInfo

	videoid := md5.Sum([]byte(video.Title))
	videoidstr := string(videoid[:])

	userid := md5.Sum([]byte(video.Author.Name))
	useridstr := string(userid[:])
	now := string(time.Now().UnixNano() / 1000000) // 转毫秒

	id := useridstr + videoidstr + now
	s.DataBasePutFile(video, id)

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
