package service

import (
	"github.com/ozline/tiktok/cmd/user/dal/db"
	"github.com/ozline/tiktok/cmd/user/pack"
	"github.com/ozline/tiktok/cmd/user/rpc"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/kitex_gen/video"
	"github.com/ozline/tiktok/pkg/utils"
)

// GetUser check token and get user's info
func (s *UserService) GetUser(req *user.InfoRequest) (*user.User, error) {
	var userResp *user.User

	// 获取用户基本信息
	userModel, err := db.GetUserByID(s.ctx, req.UserId)
	userResp = pack.User(userModel)

	if err != nil {
		return nil, err
	}

	// 关注数量
	userResp.FollowCount, err = rpc.GetFollowCount(s.ctx, &follow.FollowCountRequest{UserId: userModel.Id, Token: req.Token})

	if err != nil {
		return nil, err
	}

	// 粉丝数量
	userResp.FollowerCount, err = rpc.GetFollowerCount(s.ctx, &follow.FollowerCountRequest{UserId: userModel.Id, Token: req.Token})

	if err != nil {
		return nil, err
	}

	// 是否关注
	claims, err := utils.CheckToken(req.Token)

	if err != nil {
		return nil, err
	}

	userResp.IsFollow, err = rpc.IsFollow(s.ctx, &follow.IsFollowRequest{UserId: claims.UserId, Token: req.Token, ToUserId: userModel.Id})

	// klog.Infof("current userid: %v, to userid: %v, isfollow: %v\n", claims.UserId, userModel.Id, userResp.IsFollow)

	if err != nil {
		return nil, err
	}

	// 产品数量
	userResp.WorkCount, err = rpc.GetWorkCount(s.ctx, &video.GetWorkCountRequest{UserId: userModel.Id, Token: req.Token})

	if err != nil {
		return nil, err
	}

	// 喜欢数量
	userResp.FavoritedCount, err = rpc.GetFavoriteCount(s.ctx, &interaction.UserFavoriteCountRequest{UserId: userModel.Id, Token: req.Token})

	if err != nil {
		return nil, err
	}

	// 获赞数量
	userResp.TotalFavorited, err = rpc.GetTotalFavorited(s.ctx, &interaction.UserTotalFavoritedRequest{UserId: userModel.Id, Token: req.Token})

	if err != nil {
		return nil, err
	}

	return userResp, nil
}
