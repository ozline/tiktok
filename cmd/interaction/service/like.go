package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
)

func (s *InteractionService) Like(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if exist {
		return errno.LikeAlreadyExistError
	}

	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}

	if ok {
		if err := cache.AddVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
			return err
		}
	}

	err = db.IsFavorited(s.ctx, userID, req.VideoId, 1)
	if err == nil {
		return errno.LikeAlreadyExistError
	}
	// write into mysql
	err = db.IsFavoriteExist(s.ctx, userID, req.VideoId)
	// no exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fav := &db.Favorite{
			VideoID: req.VideoId,
			UserID:  userID,
			Status:  1,
		}
		return db.FavoriteCreate(s.ctx, fav)
	}
	// not gorm.ErrRecordNotFound error
	if err != nil {
		return err
	}
	// exist
	return db.UpdateFavoriteStatus(s.ctx, userID, req.VideoId, 1)
}
