package mq

import (
	"context"
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/ozline/tiktok/cmd/chat/dal/cache"
	"github.com/ozline/tiktok/cmd/chat/dal/db"

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
		queueName: queueName, // friendQue groupQue
	}

	ch, err := ChatMQCli.conn.Channel()
	if err != nil {
		klog.Error(err)
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
	_, err := c.channel.QueueDeclare(
		c.queueName,
		// 是否持久化
		false,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		return err
	}
	// json.marshal 可序列化结构体为二进制byte类型
	// 然后就可以通过消息队列进行传参，
	// 在消费者方面只需要通过unmarshal进行反序列化就可以得到结构体

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
		klog.Error(err)
		return
	}
	msg, err := r.channel.Consume(
		r.queueName,
		// 用来区分多个消费者
		"",
		// 是否自动应答
		true,
		// 是否具有排他性
		false,
		// 如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		// 消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		klog.Error(err)
		return
	}
	klog.Info("[*] Waiting for messages,To exit press CTRL+C")
	go r.DealWithMessageToUser(msg)
	forever := make(chan bool)
	<-forever
}
func (c *ChatMQ) DealWithMessageToUser(msg <-chan amqp.Delivery) {
	for req := range msg {
		middle_message := new(MiddleMessage)
		err := sonic.Unmarshal(req.Body, middle_message)
		if err != nil {
			klog.Error(err)
			continue
		}
		message := new(cache.Message)
		err = convertForMysql(message, middle_message)
		if err != nil {
			klog.Error(err)
			continue
		}
		Mu.Lock()
		err = db.DB.Create(&message).Error
		if err != nil {
			klog.Error(err)
			Mu.Unlock()
			continue
		}

		key := strconv.FormatInt(message.FromUserId, 10) + "-" + strconv.FormatInt(message.ToUserId, 10)
		revkey := strconv.FormatInt(message.ToUserId, 10) + "-" + strconv.FormatInt(message.FromUserId, 10)
		middle_message.IsReadNum = append(middle_message.IsReadNum, middle_message.FromUserId)
		msg, err := sonic.Marshal(middle_message)
		if err != nil {
			klog.Error(err)
			Mu.Unlock()
			continue
		}
		err = cache.MessageInsert(context.TODO(), key, revkey, message.CreatedAt.UnixMilli(), string(msg))
		if err != nil {
			klog.Error(err)
			Mu.Unlock()
			continue
		}
		Mu.Unlock()
	}
}

func convertForMysql(message *cache.Message, tempMessage *MiddleMessage) (err error) {
	message.Id = tempMessage.Id
	message.ToUserId = tempMessage.ToUserId
	message.FromUserId = tempMessage.FromUserId
	message.Content = tempMessage.Content
	message.CreatedAt, err = time.ParseInLocation(time.RFC3339, tempMessage.CreatedAt, time.Local)
	if err != nil {
		return err
	}
	return
}
