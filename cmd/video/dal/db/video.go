package db

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	PlayUrl   string `json:"play_url"`
	CoverUrl  string `json:"cover_url"`
	Title     string `json:"title"`
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
	if err := DB.WithContext(ctx).Where("id IN ?", videoId).Order("created_at DESC").Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
func GetVideoInfoByTime(ctx context.Context, latestTime string) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("created_at < ?", latestTime).Order("created_at DESC").Limit(30).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
func GetVideoInfoByUid(ctx context.Context, uid int64) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("user_id = ?", uid).Order("created_at DESC").Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
func GetWorkCountByUid(ctx context.Context, uid int64) (workCount int64, err error) {
	if err = DB.WithContext(ctx).Where("user_id = ?", uid).Count(&workCount).Error; err != nil {
		return 0, err
	}
	return
}
func GetVideoIDByUid(ctx context.Context, uid int64) (videoIDList []int64, err error) {
	if err = DB.WithContext(ctx).Select("id").Where("user_id = ?", uid).Find(&videoIDList).Error; err != nil {
		return nil, err
	}
	return
}
