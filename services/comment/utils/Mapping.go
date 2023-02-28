package utils

import (
	"github.com/ozline/tiktok/kitex_gen/tiktok/comment"
	"github.com/ozline/tiktok/services/comment/model"
)

func CommentMapping(input []*model.Comment) []*comment.Comment {
	comments := make([]*comment.Comment, len(input))
	for i, v := range input {
		comments[i] = &comment.Comment{
			Id:            v.ID,
			Uid:           v.UserID,
			Vid:           v.VideoID,
			Ctime:         v.CreatedAt.Format("01-02"),
			Content:       v.Content,
			LikeCount:     v.LikeCount,
			IsUploderLike: v.IsUploaderLike,
		}
	}
	return comments
}

func FavoriteMapping(input []*model.VideoFavorite) []int64 {
	favorites := make([]int64, len(input))

	for i, v := range input {
		favorites[i] = v.VideoID
	}

	return favorites
}
