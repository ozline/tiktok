package mq

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/follow/dal/cache"
	"github.com/ozline/tiktok/cmd/follow/dal/db"
	"github.com/ozline/tiktok/kitex_gen/follow"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/errno"
	"github.com/ozline/tiktok/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (f *FollowMQ) Publish(ctx context.Context, body string) error {
	_, err := f.channel.QueueDeclare(
		f.queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		klog.Error(err)
		return err
	}

	err = f.channel.PublishWithContext(ctx,
		f.exchange,
		f.queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // 持久化模式
			ContentType:  "application/json",
			Body:         []byte(body),
		},
	)
	if err != nil {
		klog.Error(err)
		return err
	}

	klog.Info("send msg success")
	return nil
}

func (f *FollowMQ) Consume(ctx context.Context) (<-chan amqp.Delivery, error) {
	_, err := f.channel.QueueDeclare(
		f.queueName,
		true, // 持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	msgs, err := f.channel.Consume(
		f.queueName,
		"",
		false, // 关闭自动应答
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	klog.Info("consume msg success")
	return msgs, nil
}

func (f *FollowMQ) FollowActionInsert(ctx context.Context, msg amqp.Delivery) error {
	actionMQ := new(follow.ActionRequest)
	err := sonic.Unmarshal(msg.Body, actionMQ)
	if err != nil {
		klog.Error(err)
		return err
	}

	err = action(ctx, actionMQ)
	if err != nil {
		klog.Error(err)
		return err
	}

	return nil
}

func action(ctx context.Context, req *follow.ActionRequest) error {
	claim, err := utils.CheckToken(req.Token)
	if err != nil {
		return errno.AuthorizationFailedError
	}
	action := &db.Follow{
		UserID:   claim.UserId,
		ToUserID: req.ToUserId,
	}

	Mu.Lock()
	switch req.ActionType {
	case constants.FollowAction:
		// 数据写入redis
		if err := cache.FollowAction(ctx, action.UserID, action.ToUserID); err != nil {
			klog.Error(err)
			Mu.Unlock()
			return err
		}
		// 数据写入db/更改db数据
		if err = db.FollowAction(ctx, action); err != nil {
			klog.Error(err)
			Mu.Unlock()
			return err
		}
	case constants.UnFollowAction:
		// 更改db数据
		if err = db.UnFollowAction(ctx, action); err != nil {
			klog.Error(err)
			Mu.Unlock()
			return err
		}
		time.Sleep(10 * time.Millisecond) // 延迟删除缓存中的数据
		// 删除redis中的数据
		if err = cache.UnFollowAction(ctx, action.UserID, action.ToUserID); err != nil {
			klog.Error(err)
			Mu.Unlock()
			return err
		}
	default:
		return errno.UnexpectedTypeError
	}

	if err != nil {
		klog.Error(err)
		Mu.Unlock()
		return err
	}

	Mu.Unlock()
	return nil
}

// TODO:sync resolve msg,
// add group,
// maybe need to move this func into a new package
func (s *SyncFollow) SyncFollowMQ(ctx context.Context) error {
	defer FollowMQCli.ReleaseRes()

	msgs, err := FollowMQCli.Consume(ctx)
	if err != nil {
		klog.Error(err)
		return err
	}

	var forever chan struct{}

	go func() {
		for msg := range msgs {
			klog.Infof("Resolve msg: %s", msg.Body)

			err := FollowMQCli.FollowActionInsert(ctx, msg)
			if err != nil {
				klog.Errorf("Insert follow action: %s", err)
				continue
			}

			msg.Ack(false)
		}
	}()

	<-forever

	return nil
}

// 释放资源,建议获取实例后配合defer使用
func (f *FollowMQ) ReleaseRes() {
	f.conn.Close()
	f.channel.Close()
}
