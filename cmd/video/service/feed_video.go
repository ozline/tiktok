package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/utils"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, []*user.User, []int64, []int64, []bool, error) {
	var videoList []db.Video
	var err error
	var claims *utils.Claims

	if claims, err = utils.CheckToken(*req.Token); err != nil {
		return nil, nil, nil, nil, nil, err
	}

	if exist, err := cache.IsExistVideoInfo(s.ctx, *req.LatestTime); exist == 1 {
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		videoList, err = cache.GetVideoList(s.ctx, *req.LatestTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
	} else {
		formattedTime := time.UnixMilli(*req.LatestTime).Format("2006-01-02 15:04:05")
		videoList, err = db.GetVideoInfoByTime(s.ctx, formattedTime)
		if err != nil {
			return nil, nil, nil, nil, nil, err
		}
		go cache.AddVideoList(s.ctx, videoList, *req.LatestTime)
	}
	// 创建错误组
	var eg errgroup.Group
	type result struct {
		userInfo      *user.User
		favoriteCount int64
		commentCount  int64
		isFavorite    bool
	}

	results := make([]result, len(videoList))

	// 并发调用获取user信息、favoriteCount、commentCount和isFavorite
	for i := 0; i < len(videoList); i++ {
		index := i // 在闭包中使用本地副本以避免竞态条件
		eg.Go(func() error {
			// 获取user信息
			userInfo, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[index].UserID,
				Token:  *req.Token,
			})
			if err != nil {
				return err
			}
			// 获取favoriteCount
			favoriteCount, err := rpc.GetVideoFavoriteCount(s.ctx, &interaction.VideoFavoritedCountRequest{
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			})
			if err != nil {
				return err
			}
			// 获取commentCount
			commentCount, err := rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err != nil {
				return err
			}
			// 获取isFavorite
			isFavorite, err := rpc.GetVideoIsFavorite(s.ctx, &interaction.InteractionServiceIsFavoriteArgs{Req: &interaction.IsFavoriteRequest{
				// UserId:  videoList[index].UserID,
				UserId:  claims.UserId,
				VideoId: videoList[index].Id,
				Token:   *req.Token,
			}})
			if err != nil {
				return err
			}
			results[index] = result{
				userInfo:      userInfo,
				favoriteCount: favoriteCount,
				commentCount:  commentCount,
				isFavorite:    isFavorite,
			}
			return nil
		})
	}

	// 等待所有goroutine完成
	if err := eg.Wait(); err != nil {
		return nil, nil, nil, nil, nil, err
	}
	var userList []*user.User
	var favoriteCountList []int64
	var commentCountList []int64
	var isFavoriteList []bool

	for _, result := range results {
		userList = append(userList, result.userInfo)
		favoriteCountList = append(favoriteCountList, result.favoriteCount)
		commentCountList = append(commentCountList, result.commentCount)
		isFavoriteList = append(isFavoriteList, result.isFavorite)
	}

	return videoList, userList, favoriteCountList, commentCountList, isFavoriteList, err
}
