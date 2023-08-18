package db

import (
	"context"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

type Favorite struct {
	Id         int64
	UserId     int64
	VideoId    int64
	ActionType int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func FavoriteCreate(ctx context.Context, fav *Favorite) error {
	fav.Id = SF.NextVal()
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).Create(fav).Error; err != nil {
		return err
	}
	return nil
}

func DisLike(ctx context.Context, favId int64) error {
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("id = ?", favId).Update("action_type", 2).Error; err != nil {
		return err
	}
	return nil
}

func GetFavoriteInfo(ctx context.Context, videoId int64, userId int64) (*Favorite, error) {
	var fav *Favorite
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", userId, videoId).Find(fav).Error; err != nil {
		return nil, err
	}
	return fav, nil
}

func GetVideosByUserId(ctx context.Context, userId int64) ([]int64, error) {
	videos := make([]int64, 0)
	var favs []Favorite
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND action_type", userId, 1).Find(&favs).Error; err != nil {
		return nil, err
	}

	for _, item := range favs {
		videos = append(videos, item.VideoId)
	}

	return videos, nil
}
