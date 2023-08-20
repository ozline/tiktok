package service

import (
	"bytes"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/errno"
)

func (s *VideoService) UploadCover(req *video.PutVideoRequest) (err error) {
	coverName, ok := s.ctx.Value("coverName").(string)
	if ok {
		err = s.bucket.PutObject(config.OSS.MainDirectory+"/"+coverName, bytes.NewReader(req.Cover))

	} else {
		err = errno.ServiceInternalError
	}
	return
}
