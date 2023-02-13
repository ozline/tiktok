package main

import (
	"context"
	"github.com/golang/glog"
	"github.com/ozline/tiktok/pkg/utils/snowflake"
	chat "github.com/ozline/tiktok/services/chat/kitex_gen/tiktok/chat"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TiktokChatServiceImpl implements the last service interface defined in the IDL.
type TiktokChatServiceImpl struct{}

type Message struct {
	ID           int64  // 消息id
	To_user_id   int64  // 该消息接收者的id
	From_user_id int64  // 该消息发送者的id
	Content      string // 消息内容
	Create_time  int64  // 消息创建时间
}

// SendChatMessage implements the TiktokChatServiceImpl interface.
func (s *TiktokChatServiceImpl) SendChatMessage(ctx context.Context, req *chat.DouyinSendMessageRequest) (resp *chat.DouyinSendMessageResponse, err error) {
	//fmt.Println("----- SendChatMessage -----")
	snow, err := snowflake.NewSnowflake(int64(0), int64(0))
	if err != nil {
		glog.Error(err)
		return
	}
	message := Message{
		ID:           snow.NextVal(),
		From_user_id: req.FromUserId,
		To_user_id:   req.ToUserId,
		Content:      req.Content,
		Create_time:  req.CreateTime,
	}
	s.send_message_mysql_handler(message)
	s.receive_message_mysql_handler(message)

	response := chat.DouyinSendMessageResponse{
		StatusCode: 1,
		StatusMsg:  "Success message",
		FromUserId: req.FromUserId,
		//ToUserId:   req.ToUserId,
		Content: req.Content,
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

	var messages []Message
	//fmt.Println("ToUserId=", req.ToUserId)
	db.Where("To_user_id == ?", req.ToUserId).Find(&messages)
	//fmt.Println("len(messages)=", len(messages))
	if len(messages) > 0 {
		for index, _ := range messages {
			db.Delete(&messages[index], messages[index].ID)
		}
		fromuserids := make([]int64, len(messages))
		touserids := make([]int64, len(messages))
		contents := make([]string, len(messages))
		createTimes := make([]int64, len(messages))
		for index, message := range messages {
			fromuserids[index] = message.From_user_id
			touserids[index] = message.To_user_id
			contents[index] = message.Content
			createTimes[index] = message.Create_time
		}

		response := chat.DouyinReceiveMessageResponse{
			StatusCode:  1,
			StatusMsg:   "Success Receive",
			FromUserIds: fromuserids,
			ToUserIds:   touserids,
			Contents:    contents,
			CreateTime:  createTimes,
		}
		return &response, nil
	} else {
		response := chat.DouyinReceiveMessageResponse{
			StatusCode: 2,
			StatusMsg:  "Don't Have Any Videos",
		}
		return &response, nil
	}

}
