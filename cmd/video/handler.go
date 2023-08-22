package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/pack"
	"github.com/ozline/tiktok/cmd/video/service"
	"github.com/ozline/tiktok/config"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}
	if req.Token == "" {
		req.Token, err = utils.CreateToken(10000)
		if err != nil {
			resp.Base = pack.BuildBaseResp(errno.ParamError)
		}
	}
	if _, err := utils.CheckToken(req.Token); err != nil {
		resp.Base = pack.BuildBaseResp(errno.AuthorizationFailedError)
		return resp, nil
	}
	videoList, userList, favoriteCountList, commentCountList, err := service.NewVideoService(ctx).FeedVideo(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.NextTime = videoList[0].CreatedAt.Unix()
	resp.VideoList = pack.VideoList(videoList, userList, favoriteCountList, commentCountList)

	return
}

func (s *VideoServiceImpl) PutVideo(stream video.VideoService_PutVideoServer) (err error) {
	resp := new(video.PutVideoResponse)
	//追加位置
	var nextPos int64
	var coverName string
	var videoName string

	for {
		req, err := stream.Recv()
		if err != nil {
			resp.Base = pack.BuildBaseResp(err)
			resp.State = 0
			err = stream.Send(resp)
			return err
		}
		if _, err := utils.CheckToken(req.Token); err != nil {
			resp.Base = pack.BuildBaseResp(err)
			resp.State = 0
			err = stream.Send(resp)
			return err
		}
		if coverName == "" {
			coverName = pack.GenerateCoverName(req.UserId)
		}
		if videoName == "" {
			videoName = pack.GenerateVideoName(req.UserId)
		}
		if !req.IsFinished { //上传一部分视频
			nextPos, err = service.NewVideoService(stream.Context()).UploadVideo(req, videoName, nextPos)
			if err != nil {
				resp.Base = pack.BuildBaseResp(err)
				resp.State = 0
				err = stream.Send(resp)
				return err
			}
			resp.Base = pack.BuildBaseResp(nil)
			resp.State = 1
			err = stream.Send(resp)
			if err != nil {
				return err
			}
		} else { //当视频全部上传完成后，开始封面的上传和持久化处理
			//上传封面
			err = service.NewVideoService(stream.Context()).UploadCover(req, coverName)
			if err != nil {
				resp.Base = pack.BuildBaseResp(err)
				resp.State = 0
				err = stream.Send(resp)
				return err
			}
			//保存到数据库
			playURL := fmt.Sprintf("%s/%s/%s", config.OSS.Endpoint, config.OSS.MainDirectory, videoName)
			coverURL := fmt.Sprintf("%s/%s/%s", config.OSS.Endpoint, config.OSS.MainDirectory, coverName)
			_, err = service.NewVideoService(stream.Context()).CreateVideo(req, playURL, coverURL)
			if err != nil {
				resp.Base = pack.BuildBaseResp(err)
				resp.State = 0
				err = stream.Send(resp)
				return err
			}
			klog.Infof("视频全部传输完成")
			resp.Base = pack.BuildBaseResp(nil)
			resp.State = 2
			err = stream.Send(resp)
			//结束循环停止接收
			break
		}
	}
	stream.Close()

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
	videoList, userList, favoriteCountList, commentCountList, err := service.NewVideoService(ctx).GetPublishVideoInfo(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoList(videoList, userList, favoriteCountList, commentCountList)

	return
}
