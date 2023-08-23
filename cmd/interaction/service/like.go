package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
)

func (s *InteractionService) Like(req *interaction.FavoriteActionRequest, userId int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userId)
	if err != nil {
		return err
	}
	if exist {
		return errno.LikeAlreadyExistError
	}

	if err := cache.AddVideoLikeCount(s.ctx, req.VideoId, userId); err != nil {
		return err
	}

	// write into mysql
	err = db.IsFavoriteExist(s.ctx, userId, req.VideoId)
	// no exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fav := &db.Favorite{
			VideoID: req.VideoId,
			UserID:  userId,
			Status:  1,
		}
		return db.FavoriteCreate(s.ctx, fav)
	}
	// not gorm.ErrRecordNotFound error
	if err != nil {
		return err
	}
	//exist
	return db.UpdateFavoriteStatus(s.ctx, userId, req.VideoId, 1)
}
