package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/services/video/kitex_gen/tiktok/video/tiktokvideoservice"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdEndpoints})

	if err != nil {
		panic(err)
	}

	c, err := tiktokvideoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}

	videoClient = c
}

// func VideoGetList(ctx context.Context, req *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
// 	resp, err := videoClient.GetList(ctx, req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, nil
// }

func VideoUpload(ctx context.Context, req *video.PutVideoRequest) (bool, error) {
	resp, err := videoClient.PutVideo(ctx, req)

	if err != nil {
		return false, err
	}

	//TODO: 完善错误

	resp.ErrState = ""

	return true, nil
}
