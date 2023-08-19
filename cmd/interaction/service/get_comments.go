package service

import (
	"encoding/json"
	"github.com/ozline/tiktok/cmd/interaction/dal/cache"
	"time"

	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) (*[]db.Comment, error) {

	var comments []db.Comment
	exist, err := cache.IsExistComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	if exist == 1 {
		rComments, err := cache.GetComments(s.ctx, req.VideoId)
		if err != nil {
			return nil, err
		}
		for _, rComment := range *rComments {
			var comment db.Comment
			err = json.Unmarshal([]byte((rComment.Member).(string)), &comment)
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

	var rComments []cache.Comment
	var dates []float64
	for _, comment := range comments {
		rComments = append(rComments, cache.Comment{Id: comment.Id, UserId: comment.UserId, Content: comment.Content})
		dates = append(dates, float64(comment.CreatedAt.Unix()))
	}

	err = cache.AddComments(s.ctx, req.VideoId, &rComments, &dates)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
