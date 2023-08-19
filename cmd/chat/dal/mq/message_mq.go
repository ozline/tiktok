package mq

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var MessageMQCli *MessageMQ

type MessageMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	exchange  string
	queueName string
}

func NewMessageMQ(queueName string) *MessageMQ {
	messageMQCli := &MessageMQ{
		RabbitMQ:  *Rmq,
		queueName: queueName, //friendQue groupQue
	}

	ch, err := messageMQCli.conn.Channel()
	if err != nil {

		return nil
	}
	messageMQCli.channel = ch
	return messageMQCli
}

func InitMessageMQ() {
	MessageMQCli = NewMessageMQ("messageQueue")
	go MessageMQCli.Consumer()
}

func (c *MessageMQ) Publish(ctx context.Context, message string) error {

	klog.Info("queueName ->:", c.queueName)
	_, err := c.channel.QueueDeclare(
		c.queueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err != nil {
		return err
	}
	//json.marshal 可序列化结构体为二进制byte类型
	//然后就可以通过消息队列进行传参，
	//在消费者方面只需要通过unmarshal进行反序列化就可以得到结构体

	err = c.channel.PublishWithContext(ctx,
		c.exchange,
		c.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageMQ) Consumer() {
	defer r.destroy()
	_, err := r.channel.QueueDeclare(r.queueName, false, false, false, false, nil)

	if err != nil {
		return
	}

	//2、接收消息
	msg, err := r.channel.Consume(
		r.queueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		return
	}
	go r.dealWithMessageToCache(msg)
	//log.Printf("[*] Waiting for messages,To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}

func (r *MessageMQ) dealWithMessageToCache(msg <-chan amqp.Delivery) {
	for req := range msg {

		klog.Info("messageMQ consuming")
		// if err != nil {
		// 	klog.Info(err)
		// 	continue
		// }
		message := make([]*MiddleMessage, 0)
		err := json.Unmarshal(req.Body, &message)
		if err != nil {
			klog.Info("json error here")
			klog.Info(err)
			continue
		}
		klog.Info("message mq——----------------------->", message)
		for _, val := range message {
			mes, _ := json.Marshal(val)
			key := strconv.FormatInt(val.FromUserId, 10) + "-" + strconv.FormatInt(val.ToUserId, 10)
			cre_time, _ := time.ParseInLocation("2006-01-02T15:04:05Z", val.CreatedAt, time.Local)
			err := cache.RedisDB.ZAdd(context.TODO(), key, redis.Z{
				Score:  float64(cre_time.Unix()),
				Member: mes,
			}).Err()
			if err != nil {
				klog.Info(err)
				continue
			}
		}
	}
}
