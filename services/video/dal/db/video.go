package db

import (
	"context"

	"github.com/ozline/tiktok/kitex_gen/tiktok/video"
)

type Video struct {
	ID       int64
	UserID   int64
	Title    string
	PlayUrl  string
	CoverUrl string
}

func CreateVideo(ctx context.Context, req *video.PublishActionResquest, playURL string, coverURL string) error {
	return DB.Create(&Video{
		ID:       Sf.NextVal(),
		Title:    req.Title,
		PlayUrl:  playURL,
		CoverUrl: coverURL,
		UserID:   req.Userid,
	}).Error
}

func GetVideoList(ctx context.Context, req *video.PublishListRequest) (resp []*video.Video, err error) {
	var videos []*Video

	resp = make([]*video.Video, 0)

	if err := DB.Where("user_id = ?", req.Userid).Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Find(&videos).Error; err != nil {
		return nil, err
	}

	for _, v := range videos {
		resp = append(resp, &video.Video{
			Id:       v.ID,
			UserId:   v.UserID,
			Title:    v.Title,
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		})
	}

	return resp, nil
}

func GetFeeds(ctx context.Context, req *video.FeedRequest) (resp []*video.Video, err error) {
	var videos []*Video

	resp = make([]*video.Video, 0)

	// Randomly select videos

	// TODO: use latest time to select videos
	if err := DB.Limit(int(req.PageSize)).Offset(int((req.PageNum - 1) * req.PageSize)).Order("rand()").Find(&videos).Error; err != nil {
		return nil, err
	}

	for _, v := range videos {
		resp = append(resp, &video.Video{
			Id:       v.ID,
			UserId:   v.UserID,
			Title:    v.Title,
			PlayUrl:  v.PlayUrl,
			CoverUrl: v.CoverUrl,
		})
	}

	return resp, nil
}

func GetInfo(ctx context.Context, req *video.GetInfoRequest) (resp *video.Video, err error) {
	info := new(Video)

	if err := DB.Where("id = ?", req.VideoId).First(&info).Error; err != nil {
		return nil, err
	}

	return &video.Video{
		Id:       info.ID,
		UserId:   info.UserID,
		Title:    info.Title,
		PlayUrl:  info.PlayUrl,
		CoverUrl: info.CoverUrl,
	}, nil
}
