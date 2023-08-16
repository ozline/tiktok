package service

import (
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

// Get Messages history list
func (c *ChatService) GetMessages(req *chat.MessageListRequest, user_id int64) ([]*db.Message, error) {
	messages, err := db.GetMessageList(c.ctx, req.ToUserId, user_id)

	if err != nil {
		return nil, err
	}
	return messages, nil
}
