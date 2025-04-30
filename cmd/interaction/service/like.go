package service

import (
	"encoding/json"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
)

func (s *InteractionService) Like(req *interaction.FavoriteActionRequest, userID int64) error {
	exist, err := cache.IsVideoLikeExist(s.ctx, req.VideoId, userID)
	if err != nil {
		return err
	}
	if exist {
		return errno.LikeAlreadyExistError
	}

	ok, _, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}

	if ok {
		if err := cache.AddVideoLikeCount(s.ctx, req.VideoId, userID); err != nil {
			return err
		}
	}

	// send like msg to mq
	like := &mq.LikeEvent{
		UserID:  userID,
		VideoID: req.VideoId,
		Status:  1,
	}
	likeBody, err := json.Marshal(like)
	if err != nil {
		return err
	}

	return mq.LikeMQ.SendMessageToMQ(s.ctx, likeBody)
}
