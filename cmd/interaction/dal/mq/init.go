package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/constants"
	"github.com/ozline/tiktok/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	mqURL string
	conn  *amqp.Connection
}

type InteractionMQ struct {
	RabbitMQ
	name    string
	channel *amqp.Channel
}

var (
	mqClient *RabbitMQ
	LikeMQ   *InteractionMQ
)

func Init() {
	mqClient = &RabbitMQ{
		mqURL: utils.GetMQUrl(),
	}

	conn, err := amqp.Dial(mqClient.mqURL)
	if err != nil {
		klog.Errorf("mq init error: %v", err)
		panic(err)
	}

	mqClient.conn = conn
	LikeMQ, err = NewInteractionMQ(constants.LikeQueueName)
	if err != nil {
		klog.Errorf("mq init error: %v", err)
		panic(err)
	}
}

func NewInteractionMQ(name string) (*InteractionMQ, error) {
	channel, err := mqClient.conn.Channel()
	if err != nil {
		return nil, err
	}

	return &InteractionMQ{
		name:    name,
		channel: channel,
	}, nil
}
