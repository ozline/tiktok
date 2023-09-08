package mq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

type FollowMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	exchange  string
	queueName string
}

var (
	Rmq         *RabbitMQ
	FollowMQCli *FollowMQ
)

func Init() {
	Rmq = &RabbitMQ{
		mqurl: utils.GetMQUrl(),
	}

	conn, err := amqp.Dial(Rmq.mqurl)

	if err != nil {
		klog.Error(err)
		return
	}

	Rmq.conn = conn

	FollowMQCli, err = FollowMQInit()

	if err != nil {
		klog.Error(err)
		return
	}
}

func FollowMQInit() (*FollowMQ, error) {
	ch, err := Rmq.conn.Channel()
	if err != nil {
		return nil, err
	}

	return &FollowMQ{
		RabbitMQ:  *Rmq,
		channel:   ch,
		queueName: "FollowQueue", // 加入consts
	}, nil
}
