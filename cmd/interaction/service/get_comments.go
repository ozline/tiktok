package service

import (
	"strconv"
	"time"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/cmd/interaction/pack"
	"github.com/ozline/tiktok/cmd/interaction/rpc"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) ([]*interaction.Comment, error) {
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
		for _, rComment := range *rComments {
			var comment db.Comment
			_, err = comment.UnmarshalMsg([]byte(rComment.Member.(string)))
			if err != nil {
				return nil, err
			}
			comment.CreatedAt = time.Unix(int64(rComment.Score), 0)
			comments = append(comments, comment)
		}
	} else {
		comments, err = db.GetCommentsByVideoID(s.ctx, req.VideoId)
		if err != nil {
			return nil, err
		}

		if len(comments) != 0 {
			err = cache.AddComments(s.ctx, key, &comments)
			if err != nil {
				return nil, err
			}
		}
	}

	eg, ctx := errgroup.WithContext(s.ctx)
	commentList := make([]*interaction.Comment, len(comments))
	for index, data := range comments {
		comment := data
		commentIndex := index
		eg.Go(func() error {
			userInfo, err := rpc.UserInfo(ctx, &user.InfoRequest{
				UserId: comment.UserId,
				Token:  req.Token,
			})
			rComment := pack.Comment(&comment, userInfo)
			commentList[commentIndex] = rComment
			return err
		})
	}

	if err = eg.Wait(); err != nil {
		return nil, err
	}

	return commentList, nil
}
