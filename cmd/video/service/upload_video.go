package service

import (
	"bytes"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/config"
)

func (s *VideoService) UploadVideo(req *video.PutVideoRequest, videoName string, nowPos int64) (nextPos int64, err error) {
	klog.Infof("received block %v:", req.GetBlockId())
	nextPos, err = s.bucket.AppendObject(config.OSS.MainDirectory+"/"+videoName, bytes.NewReader(req.VideoBlock), nowPos)
	return
}
