package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, []*user.User, []int64, []int64, error) {
	var videoList []db.Video
	var err error

	if exist, err := cache.IsExistVideoInfo(s.ctx, req.LatestTime); exist == 1 {
		if err != nil {
			return nil, nil, nil, nil, err
		}
		videoList, err = cache.GetVideoList(s.ctx, req.LatestTime)
		if err != nil {
			return nil, nil, nil, nil, err
		}
	} else {
		formattedTime := time.Unix(req.LatestTime, 0).Format("2006-01-02 15:04:05")
		videoList, err := db.GetVideoInfoByTime(s.ctx, formattedTime)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		go cache.AddVideoList(s.ctx, videoList, req.LatestTime)
	}

	// 获取user信息、favoriteCount、commentCount
	userList := make([]*user.User, len(videoList))
	for i := 0; i < len(videoList); i++ {
		userList[i], err = rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: videoList[i].UserID,
			Token:  req.Token,
		})
		if err != nil {
			return nil, nil, nil, nil, err
		}
	}

	favoriteCountList := make([]int64, len(videoList))
	for i := 0; i < len(videoList); i++ {
		favoriteCountList[i], err = rpc.GetVideoFavoriteCount(s.ctx, &interaction.VideoFavoritedCountRequest{
			VideoId: videoList[i].Id,
			Token:   req.Token,
		})
		if err != nil {
			return nil, nil, nil, nil, err
		}
	}

	commentCountList := make([]int64, len(videoList))
	for i := 0; i < len(videoList); i++ {
		commentCountList[i], err = rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
			VideoId: videoList[i].Id,
			Token:   &req.Token,
		})
		if err != nil {
			return nil, nil, nil, nil, err
		}
	}
	return videoList, userList, favoriteCountList, commentCountList, err
}
