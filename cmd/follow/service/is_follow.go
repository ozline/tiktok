package service

import (
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
)

func (s *FollowService) IsFollow(req *follow.IsFollowRequest) (bool, error) {
	// 先进入redis中判断是否有关注
	ex1, err := cache.IsFollow(s.ctx, req.UserId, req.ToUserId)

	if err != nil {
		return false, err
	}

	if ex1 {
		return true, nil
	}

	ex2, err := db.IsFollow(s.ctx, req.UserId, req.ToUserId)

	if err != nil {
		if errors.Is(err, db.RecordNotFound) {
			return false, nil
		}

		klog.Errorf("db sql meet error: %v\n", err)
		return false, err
	}

	return ex2, nil
}
