package service

import (
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
)

// Action Function for the follow/close operation
func (s *FollowService) Action(req *follow.ActionRequest) (*db.Follow, error) {
	//TODO:获取用户ID
	action := &db.Follow{
		UserId:     10000, //拿到ID再说(先占个位)
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	}

	return db.FollowAction(s.ctx, action)
}
