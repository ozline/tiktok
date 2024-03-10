package task

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/dal/mq"
	"gorm.io/gorm"
)

type SyncTask struct{}

var likeMU sync.Mutex

func RunSyncMQ() {
	task := &SyncTask{}
	ctx := context.Background()
	task.RunSyncLike(ctx)
}

func (*SyncTask) RunSyncLike(ctx context.Context) {
	defer mq.LikeMQ.Release()
	// get likeMsg from mq
	msg, err := mq.LikeMQ.ConsumeMessage(ctx)
	if err != nil {
		klog.Errorf("err: %v", err)
		return
	}

	// process likeData with mysql
	go func() {
		timer := time.NewTimer(30 * time.Minute)
		likes := make([]*mq.LikeEvent, 0, 100)

		for d := range msg {
			likeData := new(mq.LikeEvent)
			err := json.Unmarshal(d.Body, likeData)
			if err != nil {
				klog.Errorf("err: %v", err)
				continue
			}

			// add to likes slice
			likeMU.Lock()
			likes = append(likes, likeData)
			likeMU.Unlock()

			// per 30 minutes or likeEvent count achieve to 100
			select {
			case <-timer.C:
				likeMU.Lock()
				err = MQLikeEventToMysql(ctx, likes)
				if err != nil {
					klog.Errorf("err: %v", err)
					likeMU.Unlock()
					continue
				}
				likeMU.Unlock()

			default:
				if len(likes) >= 100 {
					likeMU.Lock()
					err = MQLikeEventToMysql(ctx, likes)
					if err != nil {
						klog.Errorf("err: %v", err)
						likeMU.Unlock()
						continue
					}
					likeMU.Unlock()
				}
			}
		}
	}()

	// block the goroutine
	forever := make(chan struct{})
	<-forever
}

// likeData Dropped into the mysql
func MQLikeEventToMysql(ctx context.Context, likes []*mq.LikeEvent) (err error) {
	for _, like := range likes {
		switch like.Status {
		case 1:
			err = db.IsFavorited(ctx, like.UserID, like.VideoID, like.Status)
			if err == nil {
				continue
			}
			// write into mysql
			err = db.IsFavoriteExist(ctx, like.UserID, like.VideoID)
			// no exist
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fav := &db.Favorite{
					UserID:  like.UserID,
					VideoID: like.VideoID,
					Status:  like.Status,
				}
				err = db.FavoriteCreate(ctx, fav)
				if err != nil {
					continue
				}
				continue
			}
			// not gorm.RecordNotFound error
			if err != nil {
				continue
			}
			// exist
			err = db.UpdateFavoriteStatus(ctx, like.UserID, like.VideoID, like.Status)
		case 0:
			err = db.UpdateFavoriteStatus(ctx, like.UserID, like.VideoID, like.Status)
		}
	}
	return err
}
