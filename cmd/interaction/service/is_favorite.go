package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"gorm.io/gorm"
)

func (s *InteractionService) IsFavorite(req *interaction.IsFavoriteRequest) (bool, error) {
	// read from redis
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, req.UserId)
	if err != nil {
		return exist, err
	}
	if exist {
		return exist, nil
	}
	// read from mysql
	err = db.IsFavorited(s.ctx, req.UserId, req.VideoId, 1)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
