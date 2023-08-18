package pack

import (
	"github.com/ozline/tiktok/cmd/interaction/dal/db"
	"github.com/ozline/tiktok/kitex_gen/interaction"
)

func Comment(data *db.Comment) *interaction.Comment {
	if data == nil {
		return nil
	}

	return &interaction.Comment{
		Id: data.Id,
		//User:       data.UserId,
		Content:    data.Content,
		CreateDate: data.CreatedAt.Format("2006-01-02 15:04:05"),
	}

}

func Comments(data *[]db.Comment) []*interaction.Comment {

	var comments []*interaction.Comment
	for _, item := range *data {
		comment := Comment(&item)
		comments = append(comments, comment)
	}

	return comments
}
