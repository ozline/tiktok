package mq

import (
    "github.com/streadway/amqp"
    "github.com/ozline/tiktok/pkg/constants"
    "fmt"
)

type RabbitMQ struct{
    conn *amqp.Connection
    mqurl string
}

var Rmq *RabbitMQ

func InitRabbitMQ() {

	Rmq = &RabbitMQ{
		mqurl: constants.MQurl,
	}
	dial, err := amqp.Dial(Rmq.mqurl)
	if err != nil {
        fmt.Println(err)
        return
	}
	Rmq.conn = dial
	return
}

func (r *RabbitMQ) destroy(){
    r.conn.Close()
}