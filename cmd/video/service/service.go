package service

import (
	"context"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/config"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

// NewVideoService new VideoService
func NewVideoService(ctx context.Context) *VideoService {
	// bucket init
	if config.OSS == nil {
		return &VideoService{ctx: ctx, bucket: nil}
	}
	client, err := oss.New(config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret, oss.UseCname(true))
	if err != nil {
		klog.Fatal(err)
	}
	bucket, err := client.Bucket(config.OSS.BucketName)
	if err != nil {
		klog.Fatal(err)
	}
	return &VideoService{ctx: ctx, bucket: bucket}
}
