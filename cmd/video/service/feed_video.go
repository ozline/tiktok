package service

import (
	"time"

	"github.com/ozline/tiktok/cmd/video/dal/cache"
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func (s *VideoService) FeedVideo(req *video.FeedRequest) ([]db.Video, []*user.User, error) {

	if exist, err := cache.IsExistVideoInfo(s.ctx, req.LatestTime); exist == 1 {
		if err != nil {
			return nil, nil, err
		} else {
			videoList, err := cache.GetVideoList(s.ctx, req.LatestTime)
			//获取user信息
			userList := make([]*user.User, len(videoList))
			for i := 0; i < len(videoList); i++ {
				userList[i], err = rpc.GetUser(s.ctx, &user.InfoRequest{
					UserId: videoList[i].UserID,
					Token:  req.Token,
				})
				if err != nil {
					return nil, nil, err
				}
			}
			return videoList, userList, err
		}
	} else {
		formattedTime := time.Unix(req.LatestTime, 0).Format("2006-01-02 15:04:05")
		videoList, err := db.GetVideoInfoByTime(s.ctx, formattedTime)
		//获取user信息
		userList := make([]*user.User, len(videoList))
		for i := 0; i < len(videoList); i++ {
			userList[i], err = rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: videoList[i].UserID,
				Token:  req.Token,
			})
			if err != nil {
				return nil, nil, err
			}
		}
		go cache.AddVideoList(s.ctx, videoList, req.LatestTime)
		return videoList, userList, err
	}

}
