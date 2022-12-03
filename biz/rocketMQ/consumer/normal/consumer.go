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
	// 1. 创建推动消费者(Push consumer)
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("consumer-group"),
	)

	// 2. 订阅topic
	err = c.Subscribe("topic-demo", consumer.MessageSelector{},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for i := range ext {
				fmt.Printf("subscribe callback:%v \n\n", ext[i])
			}
			return consumer.ConsumeSuccess, nil
		})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 3. 启动消费
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// 4. 关闭消费者
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error:%s", err.Error())
	}

}
