package service

import (
	"fmt"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]*video.Video, error) {
	// read from redis
	videoIDList, err := cache.GetUserFavoriteVideos(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(videoIDList) != 0 {
		return nil, nil
	}

	// read from mysql
	videoIDList, err = db.GetVideosByUserId(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	err = cache.UpdateFavoriteVideoList(s.ctx, req.UserId, videoIDList)
	if err != nil {
		return nil, fmt.Errorf("err: %w", err)
	}

	return rpc.GetFavoriteVideoList(s.ctx, &video.GetFavoriteVideoInfoRequest{
		VideoId: videoIDList,
		Token:   req.Token,
	})
}
