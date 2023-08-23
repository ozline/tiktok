package db

import (
	"context"
	"errors"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	Id        int64
	UserId    int64
	VideoId   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func IsFavoriteExist(ctx context.Context, userId int64, videoId int64) (bool, error) {
	var fav *Favorite
	err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userId, videoId).First(&fav).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func FavoriteCreate(ctx context.Context, fav *Favorite) error {
	fav.Id = SF.NextVal()
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).Create(fav).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFavoriteStatus(ctx context.Context, userId int64, videoId int64, status int64) error {
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userId, videoId).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func GetVideosByUserId(ctx context.Context, userId int64) ([]int64, error) {
	videos := make([]int64, 0)
	var favs []Favorite
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND status", userId, 1).Find(&favs).Error; err != nil {
		return nil, err
	}

	for _, item := range favs {
		videos = append(videos, item.VideoId)
	}

	return videos, nil
}

func GetVideoLikeCount(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("video_id = ? AND status = 1", videoId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
