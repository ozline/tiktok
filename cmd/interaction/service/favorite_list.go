package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]int64, error) {

	videosId, err := db.GetVideosByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return videosId, nil
}
