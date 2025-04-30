package mq

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	amqp "github.com/rabbitmq/amqp091-go"
)

type LikeEvent struct {
	UserID  int64
	VideoID int64
	Status  int64
}

func (mq *InteractionMQ) Release() {
	mq.channel.Close()
	mq.conn.Close()
}

func (mq *InteractionMQ) SendMessageToMQ(ctx context.Context, body []byte) error {
	queue, err := mq.channel.QueueDeclare(
		mq.name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		klog.Errorf("declare like queue err: %v", err)
		return err
	}

	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	}

	// publish the message
	err = mq.channel.PublishWithContext(ctx,
		"",
		queue.Name,
		false,
		false,
		msg)
	if err != nil {
		klog.Errorf("failed to send msg : %v", err)
		return err
	}

	klog.Infof("send msg succeed: %v", body)
	return nil
}

func (mq *InteractionMQ) ConsumeMessage(ctx context.Context) (<-chan amqp.Delivery, error) {
	queue, err := mq.channel.QueueDeclare(
		mq.name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		klog.Errorf("declare like queue err: %v", err)
		return nil, err
	}

	msg, err := mq.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		klog.Errorf("failed to register a consumer : %v", err)
		return nil, err
	}
	// TODO: put likeData into mysql with goroutine

	klog.Infof("consume msg succeed: %v", msg)
	return msg, nil
}
