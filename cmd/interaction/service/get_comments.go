package service

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"strconv"
	"time"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) (*[]db.Comment, error) {
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}

	var comments []db.Comment
	exist, err := cache.IsExistComment(s.ctx, videoId)
	if err != nil {
		return nil, err
	}
	if exist == 1 {
		rComments, err := cache.GetComments(s.ctx, videoId)
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

	comments, err = db.GetCommentsByVideoID(s.ctx, videoId)
	if err != nil {
		return nil, err
	}
	if len(comments) != 0 {
		err = cache.AddComments(s.ctx, videoId, &comments)
		if err != nil {
			return nil, err
		}
	}
	return &comments, nil
}
