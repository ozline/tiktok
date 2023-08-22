package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/redis/go-redis/v9"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]int64, error) {
	// read from redis
	videoIdList, err := cache.GetUserFavoriteVideos(s.ctx, req.UserId)
	if err == redis.Nil {
		return nil, errors.New("you have no favorite video")
	}
	if err != nil {
		return nil, err
	}
	if len(videoIdList) != 0 {
		return videoIdList, nil
	}
	// read from mysql
	videoIdList, err = db.GetVideosByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	err = cache.UpdateFavoriteVideoList(s.ctx, req.UserId, videoIdList)
	if err != nil {
		return videoIdList, errors.New("update cache fail")
	}
	return videoIdList, nil
}
