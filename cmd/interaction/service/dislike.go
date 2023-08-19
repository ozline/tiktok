package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userId int64) error {
	if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userId); err != nil {
		return err
	}
	// TODO: write into mysql periodically

	fav, err := db.GetFavoriteInfo(s.ctx, req.VideoId, userId)
	if err != nil {
		return err
	}

	if err := db.DisLike(s.ctx, fav.Id); err != nil {
		return err
	}

	return nil
}
