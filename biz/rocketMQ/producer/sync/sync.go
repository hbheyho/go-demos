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
	// 1. 生成一个producer
	p, _ := rocketmq.NewProducer(
		// 设置nameSrvAddr - 生产者或消费者能够通过名称服务器(Name Server)查找各个topic相应的Broker IP列表
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2), // 发送失败超时时间
		// 同一类Producer的集合，这类Producer发送同一类消息且发送逻辑一致
		// 如果发送的是事务消息且原始生产者在发送之后崩溃，则Broker服务器会联系同一生产者组的其他生产者实例以提交或回溯消费
		producer.WithGroupName("producer-group"),
	)

	// 2. 开始连接
	err := p.Start()
	if err != nil {
		fmt.Printf("failed to start producer: %s\n", err.Error())
		os.Exit(1)
	}

	// 3. 消息发送
	// 同步消息：producer向 broker 发送消息, 执行 API 时同步等待, 直到broker 服务器返回发送结果
	topic := "topic-demo"
	for i := 0; i < 100; i++ {
		message := &primitive.Message{
			Topic: topic,
			Body:  []byte("Hello RocketMQ Go client" + strconv.Itoa(i)),
		}
		response, err := p.SendSync(context.Background(), message) // 同步发送
		if err != nil {
			fmt.Printf("send message error:%s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", response.String())
		}
	}

	// 4. 关闭生产者
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error:%s", err.Error())
	}
}
