package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
	"golang.org/x/sync/errgroup"
)

func (s *VideoService) GetFavoriteVideoInfo(req *video.GetFavoriteVideoInfoRequest) ([]db.Video, []*user.User, []int64, []int64, error) {
	videoList, err := db.GetVideoInfoByID(s.ctx, req.VideoId)

	// 创建错误组
	var eg errgroup.Group
	type result struct {
		userInfo      *user.User
		favoriteCount int64
		commentCount  int64
	}
	results := make([]result, len(videoList))
	// 并发调用获取user信息、favoriteCount和commentCount
	for i := 0; i < len(videoList); i++ {
		index := i // 在闭包中使用本地副本以避免竞态条件
		eg.Go(func() error {
			// 获取user信息
			userInfo, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[index].UserID,
				Token:  req.Token,
			})
			if err != nil {
				return err
			}

			// 获取favoriteCount
			favoriteCount, err := rpc.GetVideoFavoriteCount(s.ctx, &interaction.VideoFavoritedCountRequest{
				VideoId: videoList[index].Id,
				Token:   req.Token,
			})
			if err != nil {
				return err
			}

			// 获取commentCount
			commentCount, err := rpc.GetCommentCount(s.ctx, &interaction.CommentCountRequest{
				VideoId: videoList[index].Id,
				Token:   &req.Token,
			})
			if err != nil {
				return err
			}
			results[index] = result{
				userInfo:      userInfo,
				favoriteCount: favoriteCount,
				commentCount:  commentCount,
			}
			return nil
		})
	}
	// 等待所有goroutine完成
	if err := eg.Wait(); err != nil {
		return nil, nil, nil, nil, err
	}

	var userList []*user.User
	var favoriteCountList []int64
	var commentCountList []int64

	for _, result := range results {
		userList = append(userList, result.userInfo)
		favoriteCountList = append(favoriteCountList, result.favoriteCount)
		commentCountList = append(commentCountList, result.commentCount)
	}

	return videoList, userList, favoriteCountList, commentCountList, err
}
