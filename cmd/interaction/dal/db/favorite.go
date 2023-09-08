package db

import (
	"context"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	ID        int64
	UserID    int64
	VideoID   int64
	Status    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func IsFavoriteExist(ctx context.Context, userID int64, videoID int64) error {
	var fav Favorite
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userID, videoID).First(&fav).Error
}

func IsFavorited(ctx context.Context, userID int64, videoID int64, status int64) error {
	var fav Favorite
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ? AND status = ?", userID, videoID, status).First(&fav).Error
}

func FavoriteCreate(ctx context.Context, fav *Favorite) error {
	fav.ID = SF.NextVal()
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).Create(fav).Error
}

func UpdateFavoriteStatus(ctx context.Context, userID int64, videoID int64, status int64) error {
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userID, videoID).Update("status", status).Error
}

func GetVideosByUserId(ctx context.Context, userID int64) ([]int64, error) {
	var favs []Favorite
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND status = 1", userID).Find(&favs).Error; err != nil {
		return nil, err
	}

	videos := make([]int64, 0, len(favs))
	for _, fav := range favs {
		videos = append(videos, fav.VideoID)
	}

	return videos, nil
}

func GetVideoLikeCount(ctx context.Context, videoID int64) (int64, error) {
	var count int64
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("video_id = ? AND status = 1", videoID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserFavoriteCount(ctx context.Context, userID int64) (int64, error) {
	var count int64
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND status = 1", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
