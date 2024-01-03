package mq

import (
	"context"
	"sync"

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
	Mu          sync.Mutex
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

	ctx := context.Background()
	go run(ctx)
}

func FollowMQInit() (*FollowMQ, error) {
	ch, err := Rmq.conn.Channel()
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &FollowMQ{
		RabbitMQ:  *Rmq,
		channel:   ch,
		queueName: "FollowQueue", // 加入consts
	}, nil
}

func run(ctx context.Context) {
	fSync := new(SyncFollow)
	err := fSync.SyncFollowMQ(ctx)
	if err != nil {
		klog.Errorf("SyncFollowMQ:%s", err)
	}
}
