package service

import (
	"encoding/json"
	"errors"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
)

func (s *InteractionService) Dislike(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if !exist {
		err := db.IsFavorited(s.ctx, userID, req.VideoId, 1)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errno.LikeNoExistError
		}
	}

	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}
	if ok {
		if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
			return err
		}
	}

	// send dislike msg to mq
	disLike := &mq.LikeEvent{
		UserID:  userID,
		VideoID: req.VideoId,
		Status:  0,
	}
	disLikeBody, err := json.Marshal(disLike)
	if err != nil {
		return err
	}

	return mq.LikeMQ.SendMessageToMQ(s.ctx, disLikeBody)
}
