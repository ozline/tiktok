package service

import (
	"errors"

	"github.com/bytedance/sonic"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func (c *ChatService) SendMessage(req *chat.MessagePostRequest, user_id int64, create_at string) error {
	if len(req.Content) == 0 || len(req.Content) > 1000 {
		return errors.New("character limit error")
	}
	message := &mq.MiddleMessage{
		Id:         db.SF.NextVal(),
		ToUserId:   req.ToUserId,
		FromUserId: user_id,
		Content:    req.Content,
		CreatedAt:  create_at,
	}
	trans_message, err := sonic.Marshal(message)
	if err != nil {
		return err
	}
	err = mq.ChatMQCli.Publish(c.ctx, string(trans_message))
	if err != nil {
		return err
	}
	return nil
}
