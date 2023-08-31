package pack

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
	antsGR "github.com/ozline/tiktok/pkg/ants"
)

type MessageBuild struct {
	MessageEle  *chat.Message
	Messagelist []*chat.Message
}

func BuildMessage(data []*db.Message) []*chat.Message {
	if data == nil {
		return make([]*chat.Message, 0)
	}

	message := &db.MessageBuild{
		MessageList: make([]*chat.Message, 0),
	}
	for _, val := range data {
		message.MessageElem = val
		antsGR.Wg.Add(1)
		antsGR.AntsPool.Invoke(message)
	}
	antsGR.Wg.Wait()
	for _, v := range message.MessageList {
		klog.Info("pack build ==> ", *v.CreateTime)
	}
	return message.MessageList
}
