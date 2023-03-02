package main

import (
	"context"

	video "github.com/ozline/tiktok/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/services/video/pack"
	service "github.com/ozline/tiktok/services/video/service"
)

// TiktokVideoServiceImpl implements the last service interface defined in the IDL.
type TiktokVideoServiceImpl struct{}

// PublishAction implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PublishAction(ctx context.Context, req *video.PublishActionResquest) (resp *video.PublishActionResponse, err error) {
	resp = new(video.PublishActionResponse)

	err = service.NewVideoService(ctx).PublishAction(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// PublishList implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) PublishList(ctx context.Context, req *video.PublishListRequest) (resp *video.PublishListResponse, err error) {
	resp = new(video.PublishListResponse)

	res, err := service.NewVideoService(ctx).PublishList(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.VideoList = res
	return resp, nil
}

// Feed implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)

	res, err := service.NewVideoService(ctx).GetFeeds(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.VideoList = res
	return
}

// GetInfo implements the TiktokVideoServiceImpl interface.
func (s *TiktokVideoServiceImpl) GetInfo(ctx context.Context, req *video.GetInfoRequest) (resp *video.GetInfoResponse, err error) {
	resp = new(video.GetInfoResponse)

	res, err := service.NewVideoService(ctx).GetVideo(req)

	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(errno.Success)
	resp.Video = res
	return
}
