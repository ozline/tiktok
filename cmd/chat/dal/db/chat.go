package db

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/ozline/tiktok/pkg/errno"
	"gorm.io/gorm"
	"gorm.io/hints"
)

type Message struct {
	Id         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
type MiddleMessage struct {
	Id         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	IsRead     int
	CreatedAt  string
}
type MessageBuild struct {
	MessageElem *Message
	MessageList []*chat.Message
}
type MessageArray []*Message

func GetMessageList(ctx context.Context, to_user_id int64, from_user_id int64) ([]*Message, error) {
	messageListFormMysql := make([]*Message, 0)
	err := DB.WithContext(ctx).Clauses(hints.UseIndex("to_user_id", "from_user_id")).
		Select("id", "to_user_id", "from_user_id", "content", "created_at").
		Where("(to_user_id=? AND from_user_id =?) OR (to_user_id=? AND from_user_id =?) ", to_user_id, from_user_id, from_user_id, to_user_id).
		Order("created_at desc").
		Find(&messageListFormMysql).Error
	if err != nil {
		// add some logs
		klog.Error("err happen ==>", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.UserExistedError
		}
		return nil, err
	}
	// 回写redis --先返回信息，然后送到mq进行异步处理
	return messageListFormMysql, nil
}

func (d *DBAction) InsertMessage(message *Message) error {
	err := d.DB.Create(message).Error
	return err
}

func (array MessageArray) Len() int {
	return len(array)
}

func (array MessageArray) Less(i, j int) bool {
	return array[i].CreatedAt.Unix() > array[j].CreatedAt.Unix()
}

func (array MessageArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
func Convert(message *Message, tempMessage *MiddleMessage) (err error) {
	message.Id = tempMessage.Id
	message.ToUserId = tempMessage.ToUserId
	message.FromUserId = tempMessage.FromUserId
	message.Content = tempMessage.Content
	message.CreatedAt, err = time.Parse(time.RFC3339, tempMessage.CreatedAt)
	if err != nil {
		return err
	}
	return
}
