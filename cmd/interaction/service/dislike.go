package service

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userId int64) error {
	if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId); err != nil {
		return err
	}
	// TODO: write into mysql periodically
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return err
	}

	fav, err := db.GetFavoriteInfo(s.ctx, videoId, userId)
	if err != nil {
		return err
	}

	if err := db.DisLike(s.ctx, fav.Id); err != nil {
		return err
	}

	return nil
}
