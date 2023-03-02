package service

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/pkg/constants"
	utils "github.com/ozline/tiktok/pkg/utils/encode"
	OSS "github.com/ozline/tiktok/pkg/utils/oss"
	"github.com/ozline/tiktok/services/video/dal/db"
)

type VideoService struct {
	ctx context.Context
	oss *OSS.OSS
}

func NewVideoService(ctx context.Context) *VideoService {
	return &VideoService{ctx: ctx}
}

func (vs *VideoService) EnableOSS() error {
	res, err := OSS.NewOSS(constants.OSSEndpoint, constants.OSSAccessKeyID, constants.OSSAccessKeySecret, constants.MainDirectory)

	if err != nil {
		return err
	}

	vs.oss = res
	return nil
}

func (vs *VideoService) PublishAction(req *video.PublishActionResquest) error {
	filename := utils.MD5Bytes(req.Data) + ".mp4" // Generate a unique file name

	err := vs.EnableOSS()

	if err != nil {
		return err
	}

	playurl, err := vs.oss.UploadObject(filename, req.Data) // Upload video to OSS

	if err != nil {
		return err
	}

	// TODO: Use FFMPEG to get coverURL

	err = db.CreateVideo(vs.ctx, req, playurl, playurl)

	if err != nil {
		return err
	}

	return nil
}

func (vs *VideoService) PublishList(req *video.PublishListRequest) ([]*video.Video, error) {

	videos, err := db.GetVideoList(vs.ctx, req)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (vs *VideoService) GetFeeds(req *video.FeedRequest) ([]*video.Video, error) {

	feeds, err := db.GetFeeds(vs.ctx, req)

	if err != nil {
		return nil, err
	}

	// TODO: Add Nexttime func

	return feeds, nil
}

func (vs *VideoService) GetVideo(req *video.GetInfoRequest) (*video.Video, error) {

	video, err := db.GetInfo(vs.ctx, req)

	if err != nil {
		return nil, err
	}

	return video, nil
}
