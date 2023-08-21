package service

import (
	"github.com/ozline/tiktok/cmd/video/dal/db"
	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/cmd/video/rpc"
	"github.com/ozline/tiktok/kitex_gen/user"
)

func (s *VideoService) GetPublishVideoInfo(req *video.GetPublishListRequest) ([]db.Video, []*user.User, error) {
	videoList, err := db.GetVideoInfoByUid(s.ctx, req.UserId)
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
