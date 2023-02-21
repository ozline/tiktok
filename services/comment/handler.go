package main

import (
	"context"
	"errors"

	comment "github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/comment/model"
	"github.com/ozline/tiktok/services/comment/service"
	"github.com/ozline/tiktok/services/comment/utils"
)

// TiktokCommentServiceImpl implements the last service interface defined in the IDL.
type TiktokCommentServiceImpl struct{}

// List implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) List(ctx context.Context, req *comment.ListReq) (resp *comment.ListResp, err error) {
	return utils.ListRespBuilder(func() (resp *comment.ListResp, err error) {
		db := model.DB().Model(&model.Comment{})
		switch req.Type {
		case comment.ListType_like:
			// db = db.Joins("comment_like", db.Where(&model.CommentLike{UserID: req.Uid}))
			return nil, errors.New("not implement")
		case comment.ListType_video:
			db = db.Where(&model.Comment{VideoID: req.Vid})
		case comment.ListType_comment:
			db = db.Where(&model.Comment{UserID: req.Uid})
		default:
			return nil, errors.New("unexpected type")
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
		return &comment.PostResp{ContendId: id},
			model.DB().Create(&model.Comment{
				ID:             id,
				UserID:         req.Uid,
				VideoID:        req.Vid,
				Content:        req.Content,
				IsUploaderLike: true,
			}).Error
	})
}

// SetLike implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) SetLike(ctx context.Context, req *comment.LikeReq) (resp *comment.LikeResp, err error) {
	return utils.LikeRespBuilder(func() (resp *comment.LikeResp, err error) {
		err = model.DB().Where(&model.Comment{ID: req.CommentId}).Take(&model.Comment{}).Error
		if err != nil {
			return
		}
		if req.IsLike {
			return &comment.LikeResp{IsLike: true}, service.CreateLike(req.Uid, req.CommentId)
		}
		return &comment.LikeResp{IsLike: false}, service.DeleteLike(req.Uid, req.CommentId)
	})
}

// GetLike implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) GetLike(ctx context.Context, req *comment.LikeReq) (resp *comment.LikeResp, err error) {
	return utils.LikeRespBuilder(func() (resp *comment.LikeResp, err error) {
		var count int64
		err = model.DB().Model(&model.CommentLike{}).Where(&model.CommentLike{UserID: req.Uid, CommentId: req.CommentId}).Count(&count).Error
		if err != nil {
			return
		}
		return &comment.LikeResp{IsLike: count == 1}, nil
	})
}
