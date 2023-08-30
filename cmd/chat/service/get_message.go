package service

import (
	"sort"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/chat"
)

// Get Messages history list
func (c *ChatService) GetMessages(req *chat.MessageListRequest, user_id int64) ([]*db.Message, error) {
	messageList := make(db.MessageArray, 0)
	// redis  ZSET
	// RedisDB.WithContext(ctx)
	key := strconv.FormatInt(req.ToUserId, 10) + "-" + strconv.FormatInt(user_id, 10)
	revkey := strconv.FormatInt(user_id, 10) + "-" + strconv.FormatInt(req.ToUserId, 10)
	if ok := cache.MessageExist(c.ctx, key); ok != 0 {
		// 查询 a->b的消息
		mem, err := cache.MessageGet(c.ctx, key)
		if err != nil {
			klog.Info(err)
			return nil, err
		}
		// 暂时用forrange
		for _, val := range mem {
			tempMessage := new(db.MiddleMessage)
			message := new(db.Message)
			err = sonic.Unmarshal([]byte(val), &tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			err = db.Convert(message, tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			messageList = append(messageList, message)
		}
	}

	if ok := cache.MessageExist(c.ctx, revkey); ok != 0 {
		mem, err := cache.MessageGet(c.ctx, revkey)
		if err != nil {
			klog.Info(err)
			return nil, err
		}
		// 暂时用forrange
		for _, val := range mem {
			tempMessage := new(db.MiddleMessage)
			message := new(db.Message)
			err = sonic.Unmarshal([]byte(val), &tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			err = db.Convert(message, tempMessage)
			if err != nil {
				klog.Info(err)
				return nil, err
			}
			messageList = append(messageList, message)
		}
	}
	if len(messageList) > 0 {
		// 合并排序
		sort.Sort(messageList)
		return messageList, nil
	}

	messages, ok, err := db.GetMessageList(c.ctx, req.ToUserId, user_id)

	if err != nil {
		klog.Info(err)
		return nil, err
	}
	if ok {
		mq_message, err := sonic.Marshal(messages)
		if err != nil {
			klog.Info(err)
			return nil, err
		}
		err = mq.MessageMQCli.Publish(c.ctx, string(mq_message))
		if err != nil {
			klog.Info(err)
			return messages, err
		}
	}

	return messages, nil
}
