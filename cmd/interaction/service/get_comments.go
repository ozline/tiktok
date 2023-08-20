package service

import (
	"strconv"
	"time"

	"github.com/ozline/tiktok/cmd/interaction/dal/cache"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) (*[]db.Comment, error) {

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
		return &comments, nil
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

	}
	return &comments, nil
}
