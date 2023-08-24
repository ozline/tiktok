package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetUserTotalFavorited(req *interaction.UserTotalFavoritedRequest) (int64, error) {
	var total int64
	videoIDList, err := rpc.GetUserVideoList(s.ctx, &video.GetVideoIDByUidRequset{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return 0, err
	}

	for _, videoID := range videoIDList {
		// read from redis
		likeCount, err := cache.GetVideoLikeCount(s.ctx, videoID)
		if likeCount == 0 {
			// read from mysql
			likeCount, err = db.GetVideoLikeCount(s.ctx, videoID)
		}
		total += likeCount
		if err != nil {
			return total, err
		}
	}

	return total, nil
}
