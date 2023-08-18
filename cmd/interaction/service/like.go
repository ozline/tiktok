package service

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Like(req *interaction.FavoriteActionRequest, userId int64) error {
	if err := cache.AddVideoLikeCount(s.ctx, req.VideoId); err != nil {
		return err
	}
	// TODO: write into mysql periodically
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return err
	}

	fav := &db.Favorite{
		VideoId:    videoId,
		UserId:     userId,
		ActionType: 1,
	}

	if err := db.FavoriteCreate(s.ctx, fav); err != nil {
		return err
	}

	return nil
}
