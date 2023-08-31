package service

import (
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"github.com/ozline/tiktok/pkg/errno"
	"golang.org/x/sync/errgroup"
)

// DeleteComment delete comment
func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest, userId int64) (*interaction.Comment, error) {
	eg, ctx := errgroup.WithContext(s.ctx)
	comment, err := db.GetCommentByID(s.ctx, *req.CommentId)
	if err != nil {
		return nil, err
	}

	if comment.UserId != userId {
		return nil, errno.AuthorizationFailedError
	}

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		var err error
		comment, err = db.DeleteComment(ctx, comment)
		return err
	})

	key := strconv.FormatInt(comment.VideoId, 10)
	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		err := cache.Unlink(s.ctx, key)
		return err
	})
	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		ok, _, err := cache.GetCount(ctx, key)
		if err != nil {
			return err
		}
		if ok {
			err = cache.AddCount(ctx, -1, key)
		}
		return err
	})

	var userInfo *user.User
	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		var err error
		userInfo, err = rpc.UserInfo(s.ctx, &user.InfoRequest{
			UserId: userId,
			Token:  req.Token,
		})
		return err
	})

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return pack.Comment(comment, userInfo), nil
}
