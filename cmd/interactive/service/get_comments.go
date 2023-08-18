package service

import (
	"encoding/json"
	"github.com/ozline/tiktok/cmd/interactive/dal/cache"
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interactive"
	"strconv"
	"time"
)

func (s *CommentService) GetComments(req *interactive.CommentListRequest) (*[]db.Comment, error) {
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
			err = json.Unmarshal([]byte((rComment.Member).(string)), &comment)
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

	var rComments []cache.Comment
	var dates []float64
	for _, comment := range comments {
		rComments = append(rComments, cache.Comment{Id: comment.Id, UserId: comment.UserId, Content: comment.Content})
		dates = append(dates, float64(comment.CreatedAt.Unix()))
	}

	err = cache.AddComments(s.ctx, videoId, &rComments, &dates)
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
