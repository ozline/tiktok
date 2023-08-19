package db

import (
	"context"
	"errors"
	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        int64 `json:"id"`
	UserId    int64 `json:"user_id"`
	VideoId   int64
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {

	comment.Id = SF.NextVal()
	if err := DB.Table(constants.CommentTableName).WithContext(ctx).Create(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func DeleteComment(ctx context.Context, comment *Comment) (*Comment, error) {
	if err := DB.Table(constants.CommentTableName).WithContext(ctx).Delete(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func GetCommentByID(ctx context.Context, commentId int64) (*Comment, error) {
	commentResp := new(Comment)

	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("id = ?", commentId).First(&commentResp).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}

	return commentResp, nil
}

func GetCommentsByVideoID(ctx context.Context, videoId int64) ([]Comment, error) {
	var commentsResp []Comment

	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("video_id = ?", videoId).
		Order("created_at desc").Find(&commentsResp).Error

	if err != nil {
		return nil, err
	}

	return commentsResp, nil
}

func CountCommentsByVideoID(ctx context.Context, videoId int64) (int64, error) {
	var count int64

	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("video_id = ? and deleted_at IS NULL", videoId).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
