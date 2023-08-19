package service

import (
	"github.com/bytedance/sonic"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

// Get Messages history list
func (c *ChatService) GetMessages(req *chat.MessageListRequest, user_id int64) ([]*db.Message, error) {
	messages, ok, err := db.GetMessageList(c.ctx, req.ToUserId, user_id)

	if err != nil {
		return nil, err
	}
	if ok {
		mq_message, _ := sonic.Marshal(messages)
		err = mq.MessageMQCli.Publish(c.ctx, string(mq_message))
		if err != nil {
			return messages, err
		}
	}

	return messages, nil
}
