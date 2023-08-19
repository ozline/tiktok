package mq

import (
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
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

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

var Rmq *RabbitMQ

func InitRabbitMQ() {

	Rmq = &RabbitMQ{
		mqurl: utils.GetMQUrl(),
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	if err != nil {
		klog.Info(err)
		return
	}
	Rmq.conn = dial
}

func (r *RabbitMQ) destroy() {
	r.conn.Close()
}
