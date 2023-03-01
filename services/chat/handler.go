package main

import (
	"context"
	"time"

	"github.com/golang/glog"
	chat "github.com/ozline/tiktok/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	"github.com/ozline/tiktok/services/chat/model"
	"github.com/ozline/tiktok/services/chat/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TiktokChatServiceImpl implements the last service interface defined in the IDL.
type TiktokChatServiceImpl struct{}

// SendChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) SendChatMessage(ctx context.Context, req *chat.SendMessageRequest) (resp *chat.SendMessageResponse, err error) {

	snow, err := snowflake.NewSnowflake(constants.SnowflakeDatacenterID, constants.SnowflakeWorkerID)
	if err != nil {
		glog.Error(err)
		return
	}
	nowTime := time.Now().Format("2006-01-02 15:04:05")

	message := model.Message{
		ID:           snow.NextVal(),
		From_user_id: req.FromUser,
		To_user_id:   req.ToUser,
		Content:      req.Content,
		Create_time:  nowTime,
	}
	dataServer := service.NewDataBaseService(ctx)
	dataServer.SendMessageMysqlHandler(message)
	dataServer.ReceiveMessageMysqlHandler(message)

	return &chat.SendMessageResponse{
		Base: &chat.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		Data: &chat.ChatMsg{
			FromUser:   req.FromUser,
			ToUser:     req.ToUser,
			Content:    req.Content,
			CreateTime: nowTime,
		},
	}, nil
}

// AcceptChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) AcceptChatMessage(ctx context.Context, req *chat.ReceiveMessageRequest) (resp *chat.ReceiveMessageResponse, err error) {

	db, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var messages []model.Message
	err = db.Where("To_user_id <> ?", req.ToUser).Find(&messages).Error

	if err != nil {
		return &chat.ReceiveMessageResponse{
			Base: &chat.BaseResp{
				Code: errno.ServiceErrorCode,
				Msg:  err.Error(),
			},
		}, nil
	}

	msgs := make([]*chat.ChatMsg, 0)

	for _, v := range messages {
		msgs = append(msgs, &chat.ChatMsg{
			FromUser:   v.From_user_id,
			ToUser:     v.To_user_id,
			Content:    v.Content,
			CreateTime: v.Create_time,
		})
	}

	return &chat.ReceiveMessageResponse{
		Base: &chat.BaseResp{
			Code: errno.SuccessCode,
			Msg:  errno.SuccessMsg,
		},
		Data: msgs,
	}, nil
}
