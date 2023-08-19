package mq

import (
	"context"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"
	"github.com/redis/go-redis/v9"

	amqp "github.com/rabbitmq/amqp091-go"
)

var ChatMQCli *ChatMQ

type ChatMQ struct {
	RabbitMQ
	channel   *amqp.Channel
	exchange  string
	queueName string
}

func NewChatMQ(queueName string) *ChatMQ {
	ChatMQCli := &ChatMQ{
		RabbitMQ:  *Rmq,
		queueName: queueName, //friendQue groupQue
	}

	ch, err := ChatMQCli.conn.Channel()
	if err != nil {

		return nil
	}
	ChatMQCli.channel = ch
	return ChatMQCli
}

func InitChatMQ() {
	ChatMQCli = NewChatMQ("chatQueue")
	go ChatMQCli.Consumer()
}

func (c *ChatMQ) Publish(ctx context.Context, message string) error {

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

func (r *ChatMQ) Consumer() {
	defer r.destroy()
	_, err := r.channel.QueueDeclare(r.queueName, false, false, false, false, nil)

	if err != nil {
		return
	}
	klog.Info("consume queuename--->", r.queueName)
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
	klog.Info("chatMQ consuming sume")
	if err != nil {
		klog.Info(err)
		return
	}
	klog.Info("chatMQ consuming sume")
	go r.DealWithMessageToUser(msg)
	//log.Printf("[*] Waiting for messages,To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
}
func (c *ChatMQ) DealWithMessageToUser(msg <-chan amqp.Delivery) {
	for req := range msg {
		//
		klog.Info("chatMQ consuming")
		middle_message := new(MiddleMessage)
		err := sonic.Unmarshal(req.Body, middle_message)
		if err != nil {
			klog.Info(err)
			continue
		}
		//先存mysql，然后存redis
		klog.Info("chatMQ consuming")
		message := new(cache.Message)
		err = convertForMysql(message, middle_message)
		if err != nil {
			klog.Info(err)
			continue
		}
		klog.Info("chatMQ consuming")

		err = db.DB.Create(&message).Error
		if err != nil {
			klog.Info(err)
			continue
		}
		key := strconv.FormatInt(message.FromUserId, 10) + "-" + strconv.FormatInt(message.ToUserId, 10)
		klog.Info("chatMQ redis consuming")
		err = cache.RedisDB.ZAdd(context.TODO(), key, redis.Z{
			Score:  float64(message.CreatedAt.Unix()),
			Member: string(req.Body),
		}).Err()
		if err != nil {
			klog.Info(err)
			continue
		}
	}
}

func convertForMysql(message *cache.Message, tempMessage *MiddleMessage) (err error) {
	message.ToUserId = tempMessage.ToUserId
	message.FromUserId = tempMessage.FromUserId
	message.Content = tempMessage.Content
	message.CreatedAt, err = time.ParseInLocation("2006-01-02T15:04:05Z", tempMessage.CreatedAt, time.Local)
	message.UpdatedAt, err = time.ParseInLocation("2006-01-02T15:04:05Z", tempMessage.UpdatedAt, time.Local)
	return
}
