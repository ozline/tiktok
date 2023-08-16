package pack

import (
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func BuildMessage(data []*db.Message) []*chat.Message {
	if data == nil {
		return nil
	}

	res := make([]*chat.Message, len(data))
	for _, val := range data {
		create_time := val.CreatedAt.Unix()
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
