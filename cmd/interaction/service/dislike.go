package service

import (
	"fmt"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("error: you did not like this video")
	}

	if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
		return err
	}

	// write into mysql
	return db.UpdateFavoriteStatus(s.ctx, userID, req.VideoId, 0)
}
