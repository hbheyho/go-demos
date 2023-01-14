package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个timeout的context
	timeoutContext, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	go handler(timeoutContext, 12500*time.Millisecond)
	select {
	case <-timeoutContext.Done():
		fmt.Println("main", timeoutContext.Err())
	}
}

// handler 定时处理
func handler(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handler", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
