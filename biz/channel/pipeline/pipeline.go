package main

import (
	"fmt"
)

// Channels也可以用于将多个goroutine连接在一起，一个Channel的输出作为下一个Channel的输入。这种串联的Channels就是所谓的管道（pipeline）。
// 下面的程序用两个channels将三个goroutine串联起来
// 1. 第一个goroutine是一个计数器;
// 2. 第二个goroutine是一个求平方的程序;
// 3. 第三个goroutine是一个打印程序.
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// counter
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	// squarer
	go func() {
		// Channel还支持close操作，用于关闭channel，随后对基于该channel的任何发送操作都将导致panic异常
		// 若接收已经closed的channel, 则返回零值
		//for {
		//	x, ok := <-naturals
		//	if !ok {
		//		break
		//	}
		//	squares <- x * x
		//}
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// printer
	//for {
	//	if x, ok := <-squares; !ok {
	//		break
	//	} else {
	//		fmt.Println(x)
	//	}
	//}
	for x := range squares {
		fmt.Println(x)
	}

}
