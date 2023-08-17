package db

import (
	"context"
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
	CreatedAt  string
	UpdatedAt  string
}
type MessageArray []*Message

func GetMessageList(ctx context.Context, to_user_id int64, from_user_id int64) ([]*Message, error) {
	messageList := make(MessageArray, 0)
	//redis  ZSET
	//RedisDB.WithContext(ctx)
	key := strconv.FormatInt(to_user_id, 10) + "-" + strconv.FormatInt(from_user_id, 10)
	revkey := strconv.FormatInt(from_user_id, 10) + "-" + strconv.FormatInt(to_user_id, 10)
	if ok, _ := cache.RedisDB.Exists(ctx, key).Result(); ok1 != 0 {
		//查询 a->b的消息
		mem, err := cache.RedisDB.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
			Min: strconv.Itoa(0),
			Max: strconv.Itoa(int(time.Now().Unix())),
		}).Result()
		if err != nil {
			return nil, err
		}
		//暂时用forrange
		for _, val := range mem {
			tempMessage := new(MiddleMessage)
			message := new(Message)
			err = json.Unmarshal([]byte(val), &tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			err = convert(message, tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			messageList = append(messageList, message)
		}
		//messageMember, _ := json.Marshal(temp)
	}
	if ok, _ := cache.RedisDB.Exists(ctx, revkey).Result(); ok2 != 0 {
		mem, err := cache.RedisDB.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
			Min: strconv.FormatInt(0, 10),
			Max: strconv.FormatInt(time.Now().Unix(), 10),
		}).Result()
		if err != nil {
			return nil, err
		}
		//暂时用forrange
		for _, val := range mem {
			tempMessage := new(MiddleMessage)
			message := new(Message)
			err = json.Unmarshal([]byte(val), &tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			err = convert(message, tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			messageList = append(messageList, message)
		}
	}
	if len(messageList) > 0 {
		//合并排序
		sort.Sort(MessageArray(messageList))
		//暂时用forrange
		for _, v := range messageList {
			messageRes := make([]*Message, 0)
			messageTmp := &Message{
				Id:         v.Id,
				ToUserId:   v.ToUserId,
				FromUserId: v.FromUserId,
				Content:    v.Content,
			}
			messageRes = append(messageRes, messageTmp)
			klog.Info(*v)
		}
		return nil, nil
	}

	//mysql
	messageListFormMysql := make([]*Message, 0)
	err := DB.WithContext(ctx).Where("(to_user_id=? AND from_user_id =?) OR (to_user_id=? AND from_user_id =?) ", to_user_id, from_user_id, from_user_id, to_user_id).Order("created_at desc").Find(&messageListFormMysql).Error
	if err != nil {
		// add some logs
		klog.Info("err happen")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	//回写redis --先返回信息，然后送到mq进行异步处理

	return messageListFormMysql, nil
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
func convert(message *Message, tempMessage *MiddleMessage) (err error) {
	message.Id = tempMessage.Id
	message.ToUserId = tempMessage.ToUserId
	message.FromUserId = tempMessage.FromUserId
	message.Content = tempMessage.Content
	message.CreatedAt, err = time.ParseInLocation("2006-01-02T15:04:05Z", tempMessage.CreatedAt, time.Local)
	message.UpdatedAt, err = time.ParseInLocation("2006-01-02T15:04:05Z", tempMessage.UpdatedAt, time.Local)
	return
}
