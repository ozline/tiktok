package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"os"
	"time"
)

func main() {
	// 设置推送消费者
	c, _ := rocketmq.NewPushConsumer(
		//消费组
		consumer.WithGroupName("testGroup"),
		// namesrv地址
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	// 必须先在 开始前
	err := c.Subscribe("Topic-test", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			fmt.Printf("subscribe callback:%v \n", ext[i])
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error:%s", err.Error())
	}
}
