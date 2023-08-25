package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
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
	err = db.IsFavoriteExist(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		exist = false
	}
	return exist, err
}
