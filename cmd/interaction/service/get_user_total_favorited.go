package service

import (
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/video"
)

func (s *InteractionService) GetUserTotalFavorited(req *interaction.UserTotalFavoritedRequest) (int64, error) {
	var total int64 = 0
	videoIDList, err := rpc.GetUserVideoList(s.ctx, &video.GetVideoIDByUidRequset{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(videoIDList))
	for _, videoID := range videoIDList {
		wg.Add(1)
		go func(videoID int64) {
			// recover panic
			defer func() {
				if e := recover(); e != nil {
					klog.Infof("recover panic :", e)
				}
				wg.Done()
			}()
			likeCount, err := cache.GetVideoLikeCount(s.ctx, videoID)
			if likeCount == 0 {
				// read from mysql
				likeCount, err = db.GetVideoLikeCount(s.ctx, videoID)
			}
			if err != nil {
				errChan <- err
			}
			total += likeCount
		}(videoID)
	}
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return total, err
		}
	}

	return total, nil
}
