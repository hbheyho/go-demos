package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
)

func main() {
	// 1. 创建producer
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
	)

	// 2. 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	// 3. 发送批量消息
	var messages []*primitive.Message
	for i := 0; i < 10; i++ {
		messages = append(messages, &primitive.Message{
			Topic: "topic-demo",
			Body:  []byte("Hello RocketMQ Go Client! num: " + strconv.Itoa(i)),
		})
	}
	// 消息发送
	res, err := p.SendSync(context.Background(), messages...)
	if err != nil {
		fmt.Printf("send message error: %s\n", err)
	} else {
		fmt.Printf("send message success: result=%s\n", res.String())
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
