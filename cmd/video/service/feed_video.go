package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, []*user.User, []int64, []int64, []bool, error) {
	var videoList []db.Video
	var err error

	if exist, err := cache.IsExistVideoInfo(s.ctx, *req.LatestTime); exist == 1 {
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		videoList, err = cache.GetVideoList(s.ctx, *req.LatestTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
	} else {
		formattedTime := time.Unix(*req.LatestTime, 0).Format("2006-01-02 15:04:05")
		videoList, err := db.GetVideoInfoByTime(s.ctx, formattedTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		go cache.AddVideoList(s.ctx, videoList, *req.LatestTime)
	}
	// 创建错误组
	var eg errgroup.Group
	// 并发调用获取user信息、favoriteCount、commentCount和isFavorite
	userResults := make(chan *user.User, len(videoList))
	favoriteCountResults := make(chan int64, len(videoList))
	commentCountResults := make(chan int64, len(videoList))
	isFavoriteResults := make(chan bool, len(videoList))
	// // 并发调用获取user信息、favoriteCount和commentCount
	for i := 0; i < len(videoList); i++ {
		index := i // 在闭包中使用本地副本以避免竞态条件
		eg.Go(func() error {
			// 获取user信息
			userInfo, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[index].UserID,
				Token:  *req.Token,
			})
			if err == nil {
				userResults <- userInfo
			} else {
				return err
			}

			// 获取favoriteCount
			favoriteCount, err := rpc.GetVideoFavoriteCount(s.ctx, &interaction.VideoFavoritedCountRequest{
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			})
			if err == nil {
				favoriteCountResults <- favoriteCount
			} else {
				return err
			}

			// 获取commentCount
			commentCount, err := rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err == nil {
				commentCountResults <- commentCount
			} else {
				return err
			}

			// 获取isFavorite
			isFavorite, err := rpc.GetVideoIsFavorite(s.ctx, &interaction.InteractionServiceIsFavoriteArgs{Req: &interaction.IsFavoriteRequest{
				UserId:  videoList[index].UserID,
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			}})
			if err == nil {
				isFavoriteResults <- isFavorite
			} else {
				return err
			}
			return nil
		})
	}
	// 等待所有goroutine完成
	if err := eg.Wait(); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// 关闭通道
	close(userResults)
	close(favoriteCountResults)
	close(commentCountResults)
	close(isFavoriteResults)
	// 从通道中提取结果
	var userList []*user.User
	for userInfo := range userResults {
		userList = append(userList, userInfo)
	}

	var favoriteCountList []int64
	for favoriteCount := range favoriteCountResults {
		favoriteCountList = append(favoriteCountList, favoriteCount)
	}

	var commentCountList []int64
	for commentCount := range commentCountResults {
		commentCountList = append(commentCountList, commentCount)
	}
	var isFavoriteList []bool
	for isFavorite := range isFavoriteResults {
		isFavoriteList = append(isFavoriteList, isFavorite)
	}
	return videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err
}
