package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/kitex_gen/video/videoservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"

	trace "github.com/kitex-contrib/tracer-opentracing"
)

func InitVideoRPC() {
	resolver, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	client, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(resolver),
		client.WithSuite(trace.NewDefaultClientSuite()),
	)

	if err != nil {
		panic(err)
	}

	videoClient = client
}

func GetFavoriteVideoList(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetFavoriteVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoList, nil
}

func GetUserVideoList(ctx context.Context, req *video.GetVideoIDByUidRequset) ([]int64, error) {
	resp, err := videoClient.GetVideoIDByUid(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoId, nil
}
