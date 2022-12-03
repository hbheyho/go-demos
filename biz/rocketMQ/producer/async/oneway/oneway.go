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

	// 2. 启动producer
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer err, err: %s", err.Error())
		os.Exit(1)
	}

	// 3. 发送单向消息
	// 单向消息：主要用在不特别关心发送结果的场景，例如日志发送
	topic := "topic-demo"
	for i := 0; i < 100; i++ {
		message := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello, RocketMR Go client send one way message, number " + strconv.Itoa(i)),
		}
		err = p.SendOneWay(context.Background(), message)
		if err != nil {
			fmt.Printf("send message error:%s\n", err)
		} else {
			fmt.Printf("send message success \n")
		}
	}

	// 4. 关闭producer
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error")
	}
}
