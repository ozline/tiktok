package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func (s *VideoService) GetFavoriteVideoInfo(req *video.GetFavoriteVideoInfoRequest) ([]db.Video, []*user.User, []int64, []int64, error) {
	videoList, err := db.GetVideoInfoByID(s.ctx, req.VideoId)

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
		favoriteCountList[i], err = rpc.GetFavoriteCount(s.ctx, &interaction.FavoriteCountRequest{
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
