package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
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
		client.WithSuite(trace.NewDefaultClientSuite()), // 设置链路追踪
	)

	if err != nil {
		panic(err)
	}

	interactionClient = c
}
func GetVideoFavoriteCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest) (favoriteCount int64, err error) {
	resp, err := interactionClient.VideoFavoritedCount(ctx, req)

	if err != nil {
		return 0, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.LikeCount, nil
}
func GetCommentCount(ctx context.Context, req *interaction.CommentCountRequest) (commentCount int64, err error) {
	resp, err := interactionClient.CommentCount(ctx, req)

	if err != nil {
		return 0, err
	}
	if resp.Base.Code != errno.SuccessCode {
		return 0, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.CommentCount, nil
}
func GetVideoIsFavorite(ctx context.Context, req *interaction.InteractionServiceIsFavoriteArgs) (isFavorite bool, err error) {
	resp, err := interactionClient.IsFavorite(ctx, req.Req)

	if err != nil {
		return false, err
	}
	if resp.Base.Code != errno.SuccessCode {
		return false, errno.NewErrNo(resp.Base.Code, *resp.Base.Msg)
	}

	return resp.IsFavorite, nil
}
