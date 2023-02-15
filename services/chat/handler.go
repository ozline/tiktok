package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	chat "github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
	"github.com/ozline/tiktok/services/chat/model"
	"github.com/ozline/tiktok/services/chat/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// TiktokChatServiceImpl implements the last service interface defined in the IDL.
type TiktokChatServiceImpl struct{}

// SendChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) SendChatMessage(ctx context.Context, req *chat.DouyinSendMessageRequest) (resp *chat.DouyinSendMessageResponse, err error) {
	//fmt.Println("----- SendChatMessage -----")
	snow, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		glog.Error(err)
		return
	}
	message := model.Message{
		ID:           snow.NextVal(),
		From_user_id: req.FromUserId,
		To_user_id:   req.ToUserId,
		Content:      req.Content,
		Create_time:  time.Now().Format("2006-01-02 15:04:05"),
	}
	dataServer := service.NewDataBaseService(ctx)
	dataServer.SendMessageMysqlHandler(message)
	dataServer.ReceiveMessageMysqlHandler(message)

	response := chat.DouyinSendMessageResponse{
		StatusCode: 1,
		StatusMsg:  "Success message",
		FromUserId: req.FromUserId,
		ToUserId:   req.ToUserId,
		Content:    req.Content,
	}
	return &response, nil
}

// AcceptChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) AcceptChatMessage(ctx context.Context, req *chat.DouyinReceiveMessageRequest) (resp *chat.DouyinReceiveMessageResponse, err error) {
	//fmt.Println("----- AcceptChatMessage -----")
	db, err := gorm.Open(sqlite.Open("receiveMessage.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var messages []model.Message
	db.Where("To_user_id <> ?", req.ToUserId).Find(&messages)
	fromuserids := make([]int64, len(messages))
	touserids := make([]int64, len(messages))
	contents := make([]string, len(messages))
	for index, message := range messages {
		fromuserids[index] = message.From_user_id
		touserids[index] = message.To_user_id
		contents[index] = message.Content
	}

	response := chat.DouyinReceiveMessageResponse{
		StatusCode:  1,
		StatusMsg:   "Success Receive",
		FromUserIds: fromuserids,
		ToUserIds:   touserids,
		Contents:    contents,
	}
	return &response, nil
}
