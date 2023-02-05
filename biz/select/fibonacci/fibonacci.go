package main

import (
	"fmt"
	"time"
)

// 注意⚠️: select 要有存有取, 同时有有消费者、生产者
func main() {
	var c, quit = make(chan int), make(chan int)
	// 启动一个goroutine去计算
	go fibonacci(c, quit)
	// 消费生产的计算结果
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 1
	time.Sleep(10 * time.Second)
}

// fibonacci 斐波那契计算
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for true {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
