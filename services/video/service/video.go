package service

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ozline/tiktok/kitex_gen/tiktok/video"
	"github.com/ozline/tiktok/pkg/constants"
	utils "github.com/ozline/tiktok/pkg/utils/encode"
	OSS "github.com/ozline/tiktok/pkg/utils/oss"
	"github.com/ozline/tiktok/pkg/utils/snapshot"
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
	filename := utils.MD5Bytes(req.Data) // Generate a unique file name

	err := vs.EnableOSS()

	if err != nil {
		return err
	}

	playurl, err := vs.oss.UploadObjectByBytes(filename+".mp4", req.Data) // Upload video to OSS

	if err != nil {
		return err
	}

	// Save Video To Local
	err = ioutil.WriteFile("./assets/"+filename+".mp4", req.Data, 0666) // Save video to local

	if err != nil {
		return err
	}

	// Get Cover Image
	coverfile, err := snapshot.GetSnapShot("./assets/"+filename+".mp4", filename) // Generate a cover image

	if err != nil {
		return err
	}

	_, coverfilename := filepath.Split(coverfile) // Get cover image name

	coverurl, err := vs.oss.UploadObjectByFile(coverfilename, coverfile) // Upload cover image to OSS

	if err != nil {
		return err
	}

	os.Remove(coverfile)
	os.Remove("./assets/" + filename + ".mp4")

	err = db.CreateVideo(vs.ctx, req, playurl, coverurl)

	if err != nil {
		return err
	}

	return nil
}

func (vs *VideoService) PublishList(req *video.PublishListRequest) ([]*video.Video, error) {
	return db.GetVideoList(vs.ctx, req)
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
	return db.GetInfo(vs.ctx, req)
}
