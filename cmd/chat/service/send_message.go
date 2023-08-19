package service

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

func (c *ChatService) SendMessage(req *chat.MessagePostRequest) error {

	//构造消息格式
	message := &mq.MiddleMessage{
		Id:         db.SF.NextVal(),
		ToUserId:   req.ToUserId,
		FromUserId: req.FromUserId,
		Content:    req.Content,
		CreatedAt:  time.Unix(*req.CreateTime, 0).Format("2006-01-02T15:04:05Z"),
		UpdatedAt:  time.Unix(*req.CreateTime, 0).Format("2006-01-02T15:04:05Z"),
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
