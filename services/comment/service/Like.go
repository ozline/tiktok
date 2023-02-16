package service

import (
	"github.com/ozline/tiktok/services/comment/model"
	"gorm.io/gorm"
)

func CreateLike(uid int64, commentId int64) (err error) {
	return model.DB().Transaction(func(tx *gorm.DB) error {
		err := model.DB().Create(&model.CommentLike{UserID: uid, CommentId: commentId}).Error
		if err != nil {
			return err
		}
		return model.DB().Model(&model.Comment{ID: commentId}).Update("like_count", gorm.Expr("like_count + ?", 1)).Error
	})
}

func DeleteLike(uid int64, commentId int64) (err error) {
	err = model.DB().Where(&model.CommentLike{UserID: uid, CommentId: commentId}).Take(&model.CommentLike{}).Error
	if err != nil {
		return
	}
	return model.DB().Transaction(func(tx *gorm.DB) error {
		err = model.DB().Delete(&model.CommentLike{UserID: uid, CommentId: commentId}).Error
		if err != nil {
			return err
		}
		return model.DB().Model(&model.Comment{ID: commentId}).Update("like_count", gorm.Expr("like_count - ?", 1)).Error
	})
}
