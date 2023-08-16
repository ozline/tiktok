package pack

import (
	"github.com/ozline/tiktok/cmd/interactive/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interactive"
)

func Comment(data *db.Comment) *interactive.Comment {
	if data == nil {
		return nil
	}

	return &interactive.Comment{
		Id: data.Id,
		//User:       data.UserId,
		Content:    data.Content,
		CreateDate: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}

}

func Comments(data *[]db.Comment) []*interactive.Comment {

	var comments []*interactive.Comment
	for _, item := range *data {
		comment := Comment(&item)
		comments = append(comments, comment)
	}

	return comments
}
