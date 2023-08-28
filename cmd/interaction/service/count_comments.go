package service

import (
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/ozline/tiktok/pkg/constants"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) CountComments(req *interaction.CommentCountRequest, times int) (count int64, err error) {
	videoId := req.VideoId

	key := strconv.FormatInt(videoId, 10)
	exist, rCount, err := cache.GetCount(s.ctx, key)
	if err != nil {
		return 0, err
	}

	if exist {
		count, err = strconv.ParseInt(rCount, 10, 64)
	} else {
		lockKey := cache.GetCountNXKey(key)
		ok, err := cache.Lock(s.ctx, lockKey)
		if err != nil {
			return 0, err
		}
		if !ok && times < constants.MaxRetryTimes {
			klog.Infof("count %v times", times+1)
			time.Sleep(constants.LockWaitTime)
			return s.CountComments(req, times+1)
		}
		count, err = db.CountCommentsByVideoID(s.ctx, videoId)
		if err != nil {
			return 0, err
		}
		err = cache.SetCount(s.ctx, key, count)
		if err != nil {
			return 0, err
		}
		if ok {
			err = cache.Delete(s.ctx, lockKey)
			if err != nil {
				return 0, err
			}
		}
	}

	if count < 0 {
		count = 0
	}
	return count, err
}
