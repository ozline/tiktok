package service

import (
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/constants"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest, times int) ([]*interaction.Comment, error) {
	var comments []db.Comment
	key := strconv.FormatInt(req.VideoId, 10)
	exist, err := cache.IsExistComment(s.ctx, key)
	if err != nil {
		return nil, err
	}

	if exist == 1 {
		rComments, err := cache.GetComments(s.ctx, key)
		if err != nil {
			return nil, err
		}
		if len(*rComments) == 1 && (*rComments)[0].Score == 0 {
			return []*interaction.Comment{}, nil
		}
		for i := 0; i < len(*rComments); i++ {
			var comment db.Comment
			_, err = comment.UnmarshalMsg([]byte((*rComments)[i].Member.(string)))
			if err != nil {
				return nil, err
			}
			comment.CreatedAt = time.Unix(int64((*rComments)[i].Score), 0)
			comments = append(comments, comment)
		}
	} else {
		lockKey := cache.GetCommentNXKey(key)
		ok, err := cache.Lock(s.ctx, lockKey)
		if err != nil {
			return nil, err
		}
		if !ok && times < constants.MaxRetryTimes {
			klog.Infof("get %v times", times+1)
			time.Sleep(constants.LockWaitTime)
			return s.GetComments(req, times+1)
		}

		comments, err = db.GetCommentsByVideoID(s.ctx, req.VideoId)
		if err != nil {
			return nil, err
		}

		if len(comments) != 0 {
			err = cache.AddComments(s.ctx, key, &comments)
			if err != nil {
				return nil, err
			}
		} else {
			klog.Infof("no comments for video %d", req.VideoId)
			err = cache.AddNoData(s.ctx, key)
			if err != nil {
				return nil, err
			}
		}

		if ok {
			err = cache.Delete(s.ctx, lockKey)
			if err != nil {
				return nil, err
			}
		}
	}

	eg, ctx := errgroup.WithContext(s.ctx)
	commentList := make([]*interaction.Comment, len(comments))
	ch := make(chan struct{}, constants.MaxGoroutines)
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		commentIndex := i
		ch <- struct{}{}
		eg.Go(func() error {
			defer func() {
				if e := recover(); e != nil {
					klog.Error(e)
				}
			}()
			userInfo, err := rpc.UserInfo(ctx, &user.InfoRequest{
				UserId: comment.UserId,
				Token:  req.Token,
			})
			rComment := pack.Comment(&comment, userInfo)
			commentList[commentIndex] = rComment
			<-ch
			return err
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return commentList, nil
}
