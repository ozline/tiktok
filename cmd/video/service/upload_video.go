package service

import (
	"bytes"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/errno"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest) (nextPos int64, err error) {
	videoName, ok := s.ctx.Value("videoName").(string)
	if !ok {
		return -1, errno.ServiceInternalError
	}
	klog.Infof("received block %v:", req.GetBlockId())
	nowPos, ok := s.ctx.Value("nextPos").(int64)
	if ok {
		nextPos, err = s.bucket.AppendObject(config.OSS.MainDirectory+"/"+videoName, bytes.NewReader(req.VideoBlock), nowPos)
	} else {
		return -1, errno.ServiceInternalError
	}
	return
}
