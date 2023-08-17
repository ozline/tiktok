package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	return
}

func (s *VideoServiceImpl) PutVideo(stream video.VideoService_PutVideoServer) (err error) {
	resp := new(video.PutVideoResponse)
	//文件名
	var videoName string
	var coverName string
	//追加位置
	var nextPos int64 = 0
	// 创建OSSClient实例。
	client, err := oss.New(config.OSS.Endpoint, config.OSS.AccessKeyID, config.OSS.AccessKeySecret)
	if err != nil {
		hanlerPutVideoError(stream, err)
		return nil
	}
	bucket, err := client.Bucket(config.OSS.BucketName)
	if err != nil {
		hanlerPutVideoError(stream, err)
		return nil
	}
	for {
		req, err := stream.Recv()
		if err != nil {
			hanlerPutVideoError(stream, err)
			return nil
		}
		if _, err := utils.CheckToken(req.Token); err != nil {
			hanlerPutVideoError(stream, err)
			return nil
		}
		if videoName == "" {
			videoName = generateVideoName(req.UserId)
		}
		if req.IsFinished {
			//上传封面
			err = bucket.PutObject(config.OSS.MainDirectory+"/"+generateCoverName(req.UserId), bytes.NewReader(req.Cover))
			if err != nil {
				hanlerPutVideoError(stream, err)
				return nil
			}
			//保存到数据库
			playUrl := fmt.Sprintf("%s/%s", config.OSS.Endpoint, videoName)
			coverUrl := fmt.Sprintf("%s/%s", config.OSS.Endpoint, coverName)

			_, err = service.NewVideoService(stream.Context()).CreateVideo(req, playUrl, coverUrl)
			if err != nil {
				hanlerPutVideoError(stream, err)
				return nil
			}
			log.Println("视频全部传输完成")
			resp.Base = pack.BuildBaseResp(nil)
			resp.State = 2
			stream.Send(resp)
			break
		}
		log.Printf("received block %v:", req.GetBlockId())
		nextPos, err = bucket.AppendObject(config.OSS.MainDirectory+"/"+videoName, bytes.NewReader(req.VideoBlock), nextPos)
		if err != nil {
			hanlerPutVideoError(stream, err)
			return nil
		}
		resp.Base = pack.BuildBaseResp(nil)
		resp.State = 1
		stream.Send(resp)
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
	videoList, err := service.NewVideoService(ctx).GetVideoInfo(req)
	if err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.VideoList = pack.VideoLikedList(videoList)
	return
}
func hanlerPutVideoError(stream video.VideoService_PutVideoServer, err error) {
	resp := new(video.PutVideoResponse)
	resp.Base = pack.BuildBaseResp(err)
	resp.State = 0
	stream.Send(resp)

}
func generateVideoName(UserId int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_video.mp4", UserId, year, month, day, hour, minute)
}
func generateCoverName(UserId int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_cover.jpg", UserId, year, month, day, hour, minute)
}
