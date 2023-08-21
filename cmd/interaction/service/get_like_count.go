package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/redis/go-redis/v9"
)

func (s *InteractionService) GetLikeCount(req *interaction.FavoriteCountRequest) (int64, error) {
	// read from redis
	likeCount, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err == redis.Nil {
		return 0, errors.New("this video has no like")
	}
	if err != nil {
		return 0, err
	}
	if likeCount == 0 {
		// read from mysql
		likeCount, err = db.GetVideoLikeCount(s.ctx, req.VideoId)
	}

	return likeCount, err
}
