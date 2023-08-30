package pack

import (
	"time"

	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func BuildMessage(data []*db.Message) []*chat.Message {
	if data == nil {
		return make([]*chat.Message, 0)
	}

	res := make([]*chat.Message, 0)
	for _, val := range data {
		create_time := val.CreatedAt.Format(time.DateTime)
		message := &chat.Message{
			Id:         val.Id,
			ToUserId:   val.ToUserId,
			FromUserId: val.FromUserId,
			Content:    val.Content,
			CreateTime: &create_time,
		}
		res = append(res, message)
	}
	return res
}
