package mq

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
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

// 释放资源,建议获取实例后配合defer使用
func (f *FollowMQ) ReleaseRes() {
	f.conn.Close()
	f.channel.Close()
}
