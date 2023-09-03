package pack

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/api/biz/model/api"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func MessageList(list []*chat.Message) []*api.Message {
	resp := make([]*api.Message, 0)

	var createtime string

	for _, data := range list {
		if data.CreateTime == nil {
			createtime = "0"
		} else {
			createtime = *data.CreateTime
		}

		klog.Infof("createtime: %v\n", createtime)
		resp = append(resp, &api.Message{
			ID:         data.Id,
			ToUserID:   data.ToUserId,
			FromUserID: data.FromUserId,
			Content:    data.Content,
			CreateTime: *data.CreateTime,
		})
	}

	return resp
}
