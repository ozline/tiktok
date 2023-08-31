package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/follow/followservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"
)

func InitFollowRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := followservice.NewClient(
		constants.FollowServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithLoadBalancer(loadbalance.NewWeightedRoundRobinBalancer()),
	)

	if err != nil {
		panic(err)
	}

	followClient = c
}

func GetFollowCount(ctx context.Context, req *follow.FollowCountRequest) (int64, error) {
	resp, err := followClient.FollowCount(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return *resp.FollowCount, nil
}

func GetFollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (int64, error) {
	resp, err := followClient.FollowerCount(ctx, req)

	if err != nil {
		return -1, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return -1, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return *resp.FollowerCount, nil
}

func IsFollow(ctx context.Context, req *follow.IsFollowRequest) (bool, error) {
	resp, err := followClient.IsFollow(ctx, req)

	if err != nil {
		return false, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return false, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.IsFollow, nil
}
