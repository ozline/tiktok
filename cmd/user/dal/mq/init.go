package mq

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Init() {
	conn, err := amqp.Dial(utils.GetMQUrl())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	body := "Hello World"

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		panic(err)
	}

	klog.Infof("Sent to rabbitmq: %s\n", body)
}
