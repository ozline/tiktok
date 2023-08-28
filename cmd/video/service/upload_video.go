package service

import (
	"bytes"

	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest, videoName string) (err error) {
	err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+videoName, bytes.NewReader(req.VideoFile))
	return
}
