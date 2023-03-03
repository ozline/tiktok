package db

import (
	"context"
	"time"

	"github.com/ozline/tiktok/kitex_gen/tiktok/chat"
)

type Message struct {
	Id          int64  `gorm:"column:id;type:bigint(20);not null;default:0;comment:'消息id'"`
	ToUserId    int64  `gorm:"column:to_user_id;type:bigint(20);not null;default:0;comment:'该消息接收者的id'"`
	FromUserId  int64  `gorm:"column:from_user_id;type:bigint(20);not null;default:0;comment:'该消息发送者的id'"`
	Content     string `gorm:"column:content;type:varchar(255);not null;default:'';comment:'消息内容'"`
	Create_time string `gorm:"column:create_time;type:varchar(255);not null;default:'';comment:'消息创建时间'"`
}

func CreateMessage(ctx context.Context, req *chat.SendMessageRequest) (resp *chat.ChatMsg, err error) {

	nowTime := time.Now().Format("2006-01-02 15:04:05") // format time
	id := Sf.NextVal()

	err = DB.Create(&Message{
		Id:          id,
		ToUserId:    req.ToUser,
		FromUserId:  req.FromUser,
		Content:     req.Content,
		Create_time: nowTime,
	}).Error

	if err != nil {
		return nil, err
	}

	return &chat.ChatMsg{
		Id:         id,
		FromUser:   req.FromUser,
		ToUser:     req.ToUser,
		Content:    req.Content,
		CreateTime: nowTime,
	}, nil
}

func GetMessageList(ctx context.Context, req *chat.ReceiveMessageRequest) (resp []*chat.ChatMsg, err error) {
	var messages []Message
	err = DB.Where("to_user_id = ? and from_user_id = ?", req.ToUser, req.FromUser).Find(&messages).Error
	if err != nil {
		return nil, err
	}

	for _, message := range messages {
		resp = append(resp, &chat.ChatMsg{
			Id:         message.Id,
			FromUser:   message.FromUserId,
			ToUser:     message.ToUserId,
			Content:    message.Content,
			CreateTime: message.Create_time,
		})
	}

	return resp, nil
}
