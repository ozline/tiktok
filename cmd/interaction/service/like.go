package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) Like(req *interaction.FavoriteActionRequest, userId int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userId)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("you already like the video")
	}

	if err := cache.AddVideoLikeCount(s.ctx, req.VideoId, userId); err != nil {
		return err
	}

	// write into mysql
	exist, err = db.IsFavoriteExist(s.ctx, userId, req.VideoId)
	if err != nil {
		return err
	}
	// no exist
	if !exist {
		fav := &db.Favorite{
			VideoId: req.VideoId,
			UserId:  userId,
			Status:  1,
		}
		return db.FavoriteCreate(s.ctx, fav)
	}
	//exist
	return db.UpdateFavoriteStatus(s.ctx, userId, req.VideoId, 1)
}
