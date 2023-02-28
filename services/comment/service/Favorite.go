package service

import (
	"github.com/ozline/tiktok/services/comment/model"
	"gorm.io/gorm"
)

func CreateFavorite(uid int64, videoID int64) (err error) {
	return model.DB().Transaction(func(tx *gorm.DB) error {
		err := model.DB().Create(&model.VideoFavorite{UserID: uid, VideoID: videoID}).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func DeleteFavorite(uid int64, videoID int64) (err error) {
	err = model.DB().Where(&model.VideoFavorite{UserID: uid, VideoID: videoID}).Take(&model.VideoFavorite{}).Error
	if err != nil {
		return
	}
	return model.DB().Transaction(func(tx *gorm.DB) error {
		err = model.DB().Delete(&model.VideoFavorite{UserID: uid, VideoID: videoID}).Error
		if err != nil {
			return err
		}
		return nil
	})
}
