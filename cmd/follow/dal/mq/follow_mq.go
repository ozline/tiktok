package mq

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (f *FollowMQ) Publish(ctx context.Context, body string) error {
	_, err := f.channel.QueueDeclare(
		f.queueName,
		true,
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
			DeliveryMode: amqp.Persistent,
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
		true,
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
		false,
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

// TODO:sync resolve msg
// type SyncFollow struct {
// }

// func (s *SyncFollow) RunTaskCreate(ctx context.Context) error {
// 	msgs, err := FollowMQCli.Consume(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	var forever chan struct{}

// 	//TODO:

// 	<-forever

// 	return nil
// }
