package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/video/pack"
	"github.com/ozline/tiktok/cmd/video/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
	"golang.org/x/sync/errgroup"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	if req.LatestTime == nil {
		currentTime := time.Now().UnixMilli()
		req.LatestTime = &currentTime
	}
	if req.Token == nil {
		req.Token = new(string)
		*req.Token, err = utils.CreateToken(10000)
		if err != nil {
			resp.Base = pack.BuildBaseResp(errno.ParamError)
		}
	}
	if _, err := utils.CheckToken(*req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err := service.NewVideoService(ctx).FeedVideo(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoList(videoList, userList, favoriteCountList, commentCountList, isFavoriteList)
	if len(videoList) > 0 {
		resp.NextTime = videoList[len(videoList)-1].CreatedAt.Unix()
	}

	return
}

// GetFavoriteVideoInfo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoInfo(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) (resp *video.GetFavoriteVideoInfoResponse, err error) {
	resp = new(video.GetFavoriteVideoInfoResponse)
	if len(req.VideoId) == 0 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}
	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	videoList, userList, favoriteCountList, commentCountList, err := service.NewVideoService(ctx).GetFavoriteVideoInfo(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoLikedList(videoList, userList, favoriteCountList, commentCountList)

	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
	resp = new(video.GetPublishListResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	if req.UserId < 10000 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err := service.NewVideoService(ctx).GetPublishVideoInfo(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoList(videoList, userList, favoriteCountList, commentCountList, isFavoriteList)

	return
}

// GetWorkCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetWorkCount(ctx context.Context, req *video.GetWorkCountRequest) (resp *video.GetWorkCountResponse, err error) {
	resp = new(video.GetWorkCountResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	if req.UserId < 10000 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	workCount, err := service.NewVideoService(ctx).GetWorkCount(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.WorkCount = workCount

	return
}

// GetGetVideoIDByUid implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoIDByUid(ctx context.Context, req *video.GetVideoIDByUidRequset) (resp *video.GetVideoIDByUidResponse, err error) {
	resp = new(video.GetVideoIDByUidResponse)

	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}

	if req.UserId < 10000 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, nil
	}

	videoIDList, err := service.NewVideoService(ctx).GetVideoIDByUid(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoId = videoIDList
	return
}

// PutVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	resp = new(video.PutVideoResponse)

	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return nil, errno.AuthorizationFailedError
	}
	videoName := pack.GenerateVideoName(claim.UserId)
	coverName := pack.GenerateCoverName(claim.UserId)
	// 创建错误组
	var eg errgroup.Group
	// 上传视频
	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadVideo(req, videoName)
		if err != nil {
			klog.Error(err)
			return errno.FileUploadError
		}
		return nil
	})
	// 截取并上传封面
	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadCover(req, coverName)
		if err != nil {
			klog.Error(err)
			return errno.FileUploadError
		}
		return nil
	})
	// 将视频数据写入数据库
	eg.Go(func() error {
		playURL := fmt.Sprintf("https://%s/%s/%s", config.OSS.Endpoint, config.OSS.MainDirectory, videoName)
		coverURL := fmt.Sprintf("https://%s/%s/%s", config.OSS.Endpoint, config.OSS.MainDirectory, coverName)
		_, err = service.NewVideoService(ctx).CreateVideo(req, playURL, coverURL)
		if err != nil {
			klog.Error(err)
		}
		return err
	})
	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	return resp, nil
}
