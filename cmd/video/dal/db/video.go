package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id        int64
	UserID    int64
	PlayUrl   string
	CoverUrl  string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	video.Id = SF.NextVal()
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}
func GetVideoInfoByID(ctx context.Context, videoId []int64) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("id IN ?", videoId).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
func GetVideoInfoByTime(ctx context.Context, latestTime string) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Limit(30).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
