package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetLikeCount(req *interaction.FavoriteCountRequest) (int64, error) {
	likeCount, err := cache.GetLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return 0, err
	}

	return likeCount, nil
}
