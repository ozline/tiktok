package rpc

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/pkg/errno"
)

func VideoGetList(ctx context.Context, req *video.PublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.PublishList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, nil
}

func VideoUpload(ctx context.Context, req *video.PublishActionResquest) error {
	resp, err := videoClient.PublishAction(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func GetFeeds(ctx context.Context, req *video.FeedRequest) ([]*video.Video, error) {
	resp, err := videoClient.Feed(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, nil
}

func GetVideoInfo(ctx context.Context, req *video.GetInfoRequest) (*video.Video, error) {
	resp, err := videoClient.GetInfo(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Video, nil
}
