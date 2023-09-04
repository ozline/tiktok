package service

import (
	"context"
	"sort"
	"strconv"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/ozline/tiktok/cmd/chat/dal/mq"
	"github.com/ozline/tiktok/kitex_gen/chat"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

// Get Messages history list
func (c *ChatService) GetMessages(req *chat.MessageListRequest, user_id int64) ([]*db.Message, error) {
	messageList := make(db.MessageArray, 0)
	// redis  ZSET
	// RedisDB.WithContext(ctx)
	key := strconv.FormatInt(req.ToUserId, 10) + "-" + strconv.FormatInt(user_id, 10)
	revkey := strconv.FormatInt(user_id, 10) + "-" + strconv.FormatInt(req.ToUserId, 10)
	no_empty := 0
	msg_array, err := cacheMessageDeal(c.ctx, key, &no_empty)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	messageList = append(messageList, msg_array...)
	rev_msg_array, err := cacheMessageDeal(c.ctx, revkey, &no_empty)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	messageList = append(messageList, rev_msg_array...)
	if len(messageList) > 0 {
		// 合并排序
		sort.Sort(messageList)
		return messageList, nil
	}
	if no_empty == 1 {
		// 没有新消息
		return nil, nil
	}
	messages, err := db.GetMessageList(c.ctx, req.ToUserId, user_id)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	mq_message, err := sonic.Marshal(messages)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	err = mq.MessageMQCli.Publish(c.ctx, string(mq_message))
	if err != nil {
		klog.Error(err)
		return messages, err
	}

	return messages, nil
}

func cacheMessageDeal(ctx context.Context, key string, isempty *int) (db.MessageArray, error) {
	msg_array := make(db.MessageArray, 0)
	if ok := cache.MessageExist(ctx, key); ok != 0 {
		// 查询 a->b的消息
		mem, err := cache.MessageGet(ctx, key)
		if err != nil {
			klog.Error(err)
			return nil, err
		}
		if len(mem) != 0 {
			*isempty = 1
		}
		var eg errgroup.Group
		for _, val := range mem {
			temp_val := val
			eg.Go(func() error {
				tempMessage := new(db.MiddleMessage)
				message := new(db.Message)
				err = sonic.Unmarshal([]byte(temp_val), &tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				if tempMessage.IsRead == 1 {
					return nil
				}
				err = db.Convert(message, tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				err = cache.RedisDB.ZRem(ctx, key, temp_val).Err()
				if err != nil {
					klog.Error(err)
					return err
				}
				tempMessage.IsRead = 1
				redis_msg, err := sonic.Marshal(tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				err = cache.RedisDB.ZAdd(ctx, key, redis.Z{
					Score:  float64(message.CreatedAt.UnixMilli()),
					Member: string(redis_msg),
				}).Err()
				if err != nil {
					klog.Error(err)
					return err
				}
				msg_array = append(msg_array, message)
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
	}
	return msg_array, nil
}
