package service

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]int64, error) {
	userId, err := strconv.ParseInt(req.UserId, 10, 64)
	if err != nil {
		return nil, err
	}
	videosId, err := db.GetVideosByUserId(s.ctx, userId)
	if err != nil {
		return nil, err
	}

	return videosId, nil
}
