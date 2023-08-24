package service

import (
	"strconv"
	"sync"

	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/user"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/pkg/errno"
)

// DeleteComment delete comment
func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest, userId int64) (*interaction.Comment, error) {
	var wg sync.WaitGroup
	comment, err := db.GetCommentByID(s.ctx, *req.CommentId)
	if err != nil {
		return nil, err
	}

	if comment.UserId != userId {
		return nil, errno.AuthorizationFailedError
	}

	errs := make([]error, 4)
	wg.Add(1)
	go func() {
		defer wg.Done()
		comment, err = db.DeleteComment(s.ctx, comment)
		errs[0] = err
	}()

	key := strconv.FormatInt(comment.VideoId, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		exist, err := cache.IsExistComment(s.ctx, key)
		errs[1] = err
		if err != nil {
			return
		}
		if exist == 1 {
			err = cache.DeleteComments(s.ctx, key)
			errs[2] = err
		}
	}()

	var userInfo *user.User
	wg.Add(1)
	go func() {
		defer wg.Done()
		userInfo, err = rpc.UserInfo(s.ctx, &user.InfoRequest{
			UserId: userId,
			Token:  req.Token,
		})
		errs[3] = err
	}()
	wg.Wait()

	for _, err = range errs {
		if err != nil {
			return nil, err
		}
	}

	return pack.Comment(comment, userInfo), nil
}
