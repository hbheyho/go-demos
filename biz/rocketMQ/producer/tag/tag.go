package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
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

	// 3. 根据tag发送请求
	tags := []string{"TagA", "TagB", "TagC"}
	for i := 0; i < 3; i++ {
		tag := tags[i%3]
		msg := &primitive.Message{
			Topic: "topic-demo",
			Body:  []byte("Hello RocketMQ Go Client"),
		}
		msg.WithTag(tag)
		res, err := p.SendSync(context.Background(), msg)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}

	// 4. 关闭producer
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}

}
