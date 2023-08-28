package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout*100),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	videoClient = c
}

func VideoFeed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, int64, error) {
	resp, err := videoClient.Feed(ctx, req)

	if err != nil {
		return nil, -1, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, resp.NextTime, nil
}

func PublishList(ctx context.Context, req *video.GetPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, nil
}

func VideoPublish(ctx context.Context, req *video.PutVideoRequest) error {
	resp, err := videoClient.PutVideo(ctx, &video.PutVideoRequest{
		VideoFile: req.VideoFile,
		Title:     req.Title,
		Token:     req.Token,
	})

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}
