package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"os"
	"strconv"
	"sync"
)

func main() {
	// 1. 创建一个producer
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2))

	// 2. 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}

	// 3. 发送异步消息
	// 异步消息: producer向 broker 发送消息时指定消息发送成功及发送异常的回调方法, 调用 API 后立即返回,
	// producer发送消息线程不阻塞, 消息发送成功或失败的回调任务在一个新的线程中执行
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		err := p.SendAsync(context.Background(),
			func(ctx context.Context, result *primitive.SendResult, err error) {
				if err != nil {
					fmt.Printf("send message error: %s\n", err)
				} else {
					fmt.Printf("send message success: result=%s\n", result.String())
				}
				wg.Done()
			}, primitive.NewMessage("topic-demo", []byte(("Hello RocketMQ Go Client")+strconv.Itoa(i))),
		)
		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		}
	}

	// 4. 等待消息发送完毕并停止
	wg.Wait()
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
