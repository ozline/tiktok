package main

import (
	"context"
	"time"

	comment "github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
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
			return nil, errno.NotImplementError
		case comment.ListType_video:
			db = db.Where(&model.Comment{VideoID: req.Vid})
		case comment.ListType_comment:
			db = db.Where(&model.Comment{UserID: req.Uid})
		default:
			return nil, errno.UnexpectedTypeError
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
		return &comment.PostResp{Comment: &comment.Comment{
				Id:      id,
				Uid:     req.Uid,
				Content: req.Content,
				Ctime:   time.Now().Format("01-02"),
			}},
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

// SetFavorite implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) SetFavorite(ctx context.Context, req *comment.FavoriteReq) (resp *comment.FavoriteResp, err error) {
	return utils.FavoriteRespBuilder(func() (resp *comment.FavoriteResp, err error) {
		if req.IsLike {
			return &comment.FavoriteResp{}, service.CreateFavorite(req.Uid, req.VideoId)
		}
		return &comment.FavoriteResp{}, service.DeleteFavorite(req.Uid, req.VideoId)
	})
}

// FavoriteList implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) FavoriteList(ctx context.Context, req *comment.FavoriteListReq) (resp *comment.FavoriteListResp, err error) {
	return utils.FavoriteListRespBuilder(func() (resp *comment.FavoriteListResp, err error) {
		db := model.DB().Model(&model.VideoFavorite{}).Where(&model.VideoFavorite{UserID: req.Uid})

		favorites := make([]*model.VideoFavorite, 0)
		var count int64
		err = db.Count(&count).Offset(int(req.PageNumber) - 1).Limit(int(req.PageSize)).Find(&favorites).Error

		if err != nil {
			return nil, err
		}

		return &comment.FavoriteListResp{Count: count, Videos: utils.FavoriteMapping(favorites)}, nil
	})
}

// GetVideoInfo implements the TiktokCommentServiceImpl interface.
func (s *TiktokCommentServiceImpl) GetVideoInfo(ctx context.Context, req *comment.GetVideoInfoReq) (resp *comment.GetVideoInfoResp, err error) {
	return utils.GetVideoInfoRespBuilder(func() (resp *comment.GetVideoInfoResp, err error) {
		resp = new(comment.GetVideoInfoResp)
		var count int64

		// Check favorite exist
		db := model.DB().Model(&model.VideoFavorite{}).Where(&model.VideoFavorite{UserID: req.Uid, VideoID: req.VideoId})

		if db.Error != nil {
			resp.IsFavorite = false
		} else {
			resp.IsFavorite = true
		}

		// Get Count
		db = model.DB().Model(&model.Comment{}).Where(&model.Comment{VideoID: req.VideoId}).Count(&count)

		if db.Error != nil {
			resp.CommentCount = 0
		} else {
			resp.CommentCount = count
		}

		// Get Favorites
		db = model.DB().Model(&model.VideoFavorite{}).Where(&model.VideoFavorite{VideoID: req.VideoId}).Count(&count)

		if db.Error != nil {
			resp.FavoriteCount = 0
		} else {
			resp.FavoriteCount = count
		}

		return resp, nil

	})
}
