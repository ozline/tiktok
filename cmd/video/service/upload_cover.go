package service

import (
	"bytes"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/config"
)

func (s *VideoService) UploadCover(req *video.PutVideoRequest, coverName string) (err error) {
	err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+coverName, bytes.NewReader(req.Cover))
	return
}
