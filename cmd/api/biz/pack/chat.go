package pack

import (
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func MessageList(list []*chat.Message) []*api.Message {
	resp := make([]*api.Message, 0)

	for _, data := range list {
		resp = append(resp, &api.Message{
			ID:         data.Id,
			ToUserID:   data.ToUserId,
			FromUserID: data.FromUserId,
			Content:    data.Content,
			// CreateTime: strconv.FormatInt(data.CreateTime, 10),
			// TODO: fix type error
		})
	}

	return resp
}
