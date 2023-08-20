package service

import (
	"context"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ozline/tiktok/config"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

// NewVideoService new VideoService
func NewVideoService(ctx context.Context) *VideoService {
	//初始化bucket
	if config.OSS == nil {
		return &VideoService{ctx: ctx, bucket: nil}
	}
	client, _ := oss.New(config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret, oss.UseCname(true))
	bucket, _ := client.Bucket(config.OSS.BucketName)

	return &VideoService{ctx: ctx, bucket: bucket}
}
