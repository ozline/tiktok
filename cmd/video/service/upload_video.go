package service

import (
	"bytes"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/config"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest, videoName string) (err error) {
	err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+videoName, bytes.NewReader(req.VideoFile))
	return
}
