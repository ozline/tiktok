package service

import (
	"errors"

	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
)

func (s *FollowService) IsFollow(req *follow.IsFollowRequest) (bool, error) {
	// 先进入redis中判断是否有关注
	ex1, err := cache.IsFollow(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return ex1, err
	}
	ex2, err := db.IsFollow(s.ctx, req.UserId, req.ToUserId)
	if errors.Is(err, db.RecordNotFound) { // 说明db中不存在
		return false, nil
	} else if err != nil {
		return ex2, err
	}
	return ex1 && ex2, nil
}
