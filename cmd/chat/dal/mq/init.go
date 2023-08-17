package mq

import (
	"fmt"

	"github.com/ozline/tiktok/pkg/utils"
	"github.com/streadway/amqp"
)

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
		fmt.Println(err)
		return
	}
	Rmq.conn = dial
}

func (r *RabbitMQ) destroy() {
	r.conn.Close()
}
