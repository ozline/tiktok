package main

import (
	"context"

	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	comment "github.com/ozline/tiktok/services/comment/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/services/comment/model"
	"github.com/ozline/tiktok/services/comment/utils"
	"gorm.io/gorm"
)

// TiktokCommentServiceImpl implements the last service interface defined in the IDL.
type TiktokCommentServiceImpl struct{}

// List implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) List(ctx context.Context, req *comment.ListReq) (resp *comment.ListResp, err error) {
	return utils.ListRespBuilder(func() (resp *comment.ListResp, err error) {
		db := model.DB()
		switch req.Type {
		case comment.ListType_like:
			db = db.Joins("CommentLike", db.Where(&model.CommonLike{UserID: req.Uid}))
		case comment.ListType_video:
			db = db.Where(&model.Comment{VideoID: req.Vid})
		case comment.ListType_comment:
			db = db.Where(&model.Comment{UserID: req.Uid})
		}
		var comments []*model.Comment
		var count int64
		err = db.Count(&count).Offset(int(req.PageNumber) - 1).Limit(int(req.PageSize)).Find(&comments).Error
		if err != nil {
			return nil, err
		}
		return &comment.ListResp{Count: count, Comments: utils.CommentMapping(comments)}, nil
	})
}

// Post implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) Post(ctx context.Context, req *comment.PostReq) (resp *comment.PostResp, err error) {
	return utils.PostRespBuilder(func() (resp *comment.PostResp, err error) {
		sf, err := snowflake.NewSnowflake(constants.SnowflakeWorkerID, constants.SnowflakeDatacenterID)
		if err != nil {
			return nil, err
		}
		id := sf.NextVal()
		err = model.DB().Create(&model.Comment{
			ID:             id,
			UserID:         req.Uid,
			VideoID:        req.Vid,
			Content:        req.Content,
			IsUploaderLike: true,
		}).Error
		if err != nil {
			return nil, err
		}
		return &comment.PostResp{}, nil
	})
}

// SetLike implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) SetLike(ctx context.Context, req *comment.LikeReq) (resp *comment.LikeResp, err error) {
	return utils.LikeRespBuilder(func() (resp *comment.LikeResp, err error) {
		var count int64
		err = model.DB().Where(&model.CommonLike{UserID: req.Uid, CommentId: req.CommentId}).Count(&count).Error
		if err != nil {
			return
		}
		if count > 1 { // TODO: test
			panic("wtf?")
		}
		if count == 0 {
			err = model.DB().Transaction(func(tx *gorm.DB) error {
				err := model.DB().Create(&model.CommonLike{UserID: req.Uid, CommentId: req.CommentId}).Error
				if err != nil {
					return err
				}
				// TODO: is the LikeCount name `like_count`?
				return model.DB().Model(&comment.Comment{}).Update("like_count", gorm.Expr("like_count - ?", 1)).Error
			})
			if err != nil {
				return
			}
		}
		return &comment.LikeResp{IsLike: true}, nil
	})
}

// GetLike implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) GetLike(ctx context.Context, req *comment.LikeReq) (resp *comment.LikeResp, err error) {
	return utils.LikeRespBuilder(func() (resp *comment.LikeResp, err error) {
		var count int64
		err = model.DB().Where(&model.CommonLike{UserID: req.Uid, CommentId: req.CommentId}).Count(&count).Error
		if err != nil {
			return
		}
		return &comment.LikeResp{IsLike: count == 1}, nil
	})
}
