package pack

import (
	"fmt"

	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

type MessageBuildArray []*chat.Message

func BuildMessage(data []*db.Message) []*chat.Message {
	if data == nil {
		return make([]*chat.Message, 0)
	}
	res := make(MessageBuildArray, 0)
	for _, val := range data {
		create_at := fmt.Sprintf("%v", val.CreatedAt.UnixMilli())
		msg := &chat.Message{
			Id:         val.Id,
			ToUserId:   val.ToUserId,
			FromUserId: val.FromUserId,
			Content:    val.Content,
			CreateTime: &create_at,
		}
		res = append(res, msg)
	}
	return res
}
