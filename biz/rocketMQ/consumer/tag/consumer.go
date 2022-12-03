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
	// 1. 创建consumer
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("consumer-group"),
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
	)

	// 2. 配置tag
	selector := consumer.MessageSelector{
		Type:       consumer.TAG,
		Expression: "TagA || TagC",
	}

	// 3. 订阅
	err := c.Subscribe("topic-demo", selector,
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			fmt.Printf("subscribe callback: %v \n", ext)
			return consumer.ConsumeSuccess, nil
		},
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 4. 开始消费
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)

	// 5. 停止consumer
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}
