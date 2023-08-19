package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/middleware"
)

func InitInteractionRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.Etcd.Addr})

	if err != nil {
		panic(err)
	}

	c, err := interactionservice.NewClient(
		constants.InteractionServiceName,
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

	interactionClient = c
}

func FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) error {
	resp, err := interactionClient.FavoriteAction(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return nil
}

func FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) ([]*interaction.Video, error) {
	resp, err := interactionClient.FavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.VideoList, nil
}

func CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (*interaction.Comment, error) {
	resp, err := interactionClient.CommentAction(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.Comment, nil
}

func CommentList(ctx context.Context, req *interaction.CommentListRequest) ([]*interaction.Comment, error) {
	resp, err := interactionClient.CommentList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.CommentList, nil
}
