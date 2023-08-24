package db

import (
	"context"
	"errors"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/gorm"
)

//go:generate msgp -io=false -tests=false
type Comment struct {
	Id        int64          `msg:"i"`
	UserId    int64          `msg:"u"`
	VideoId   int64          `msg:"-"`
	Content   string         `msg:"c"`
	CreatedAt time.Time      `msg:"-"`
	UpdatedAt time.Time      `msg:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" msg:"-"`
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
