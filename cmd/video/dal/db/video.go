package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64
	UserID        int64
	FavoriteCount int64
	CommentCount  int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	video.Id = SF.NextVal()
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}
func GetVideoInfo(ctx context.Context, videoId []int64) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("id IN ?", videoId).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
