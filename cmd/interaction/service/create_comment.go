package service

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"golang.org/x/sync/errgroup"
)

// CreateComment create comment
func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest, userId int64) (*interaction.Comment, error) {
	eg, ctx := errgroup.WithContext(s.ctx)
	commentModel := &db.Comment{
		VideoId: req.VideoId,
		UserId:  userId,
		Content: *req.CommentText,
	}

	comment := new(db.Comment)

	eg.Go(func() error {
		var err error
		comment, err = db.CreateComment(ctx, commentModel)
		return err
	})

	key := strconv.FormatInt(req.VideoId, 10)
	eg.Go(func() error {
		exist, err := cache.IsExistComment(ctx, key)
		if err != nil {
			return err
		}
		if exist == 1 {
			err = cache.DeleteComments(ctx, key)
		}
		return err
	})

	userInfo := new(user.User)

	eg.Go(func() error {
		var err error
		userInfo, err = rpc.UserInfo(ctx, &user.InfoRequest{
			UserId: userId,
			Token:  req.Token,
		})
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return pack.Comment(comment, userInfo), nil
}
