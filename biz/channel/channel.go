package main

import (
	"fmt"
	"time"
)

// Channel特点:
// 1. channel是goroutine之间的双向通信管道, 且数据发送完成后, 一定需要一个goroutine来进行接收 => 同步channel
// 2. channel可以作为参数、返回值、变量
// 3. go提供了bufferedChannel, 通过缓存channel数据, 减少goroutine之间的切换, 且不一定需要goroutine来进行接收数据 => 异步channel
// 4. 发送方通过close(channel)来告知接收方数据发送完毕, 此时接收方进行额外判断, n, ok := <- c, 进而停止接收, 或者通过range判断是否发送完成
// 5. 若不使用close关闭, 接收方的goroutine不会停止接收, 直到main线程结束
// 6. go语言的并发和channel基于 Communication Sequential Process(CSP)模型实现
// 7. 两个相同类型的channel可以使用==运算符比较，如果两个channel引用的是相同的对象，那么比较的结果为真。一个channel也可以和nil进行比较
// 8. 当一个被关闭的channel中已经发送的数据都被成功接收后，后续的接收操作将不再阻塞，它们会立即返回一个零值
// 9. 不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收
// 10. 试图重复关闭一个channel将导致panic异常，试图关闭一个nil值的channel也将导致panic异常
// 11. 只有在发送者所在的goroutine才会调用close函数，因此对一个只接收的channel调用close将是一个编译错误
// 12. 任何双向channel向单向channel变量的赋值操作都将导致该隐式转换。这里并没有反向转换的语法：也就是不能将一个类似chan<- int类型的单向型的channel转换为chan int类型的双向型的channel
// 13. 无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的
// 14. 同步channel: 一个channel只有发送者, 没有接受者会deadlock, 一个channel只有接收者, 没有发送者会deadlock
// 15. 异步channel: 当发送消息数量大于buffer, 也会发生deadlock, 不关心到底有没有接受者
func main() {
	ChannelDemo()
	BufferedChannelDemo()
}

// worker channel作为参数
func worker(id int, c chan int) {
	for {
		// 判断发送方数据是否发送完毕
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %d\n", id, n)
	}
	// 通过range判断是否发送完成
	/*for n := range c {
	    fmt.Printf("worker %d received %d \n", id, n)
	}*/
}

// createWorker channel 作为返回值
func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}()
	return c
}

func ChannelDemo() {
	// channel创建
	// 创建一个channel, 内容为int
	// var c chan int  其中c = nil
	// c := make(chan int)

	// 与单个channel通信
	// go worker(0, c)
	// c <- 1

	// 与多个worker通信 => channel作为参数
	// 为每个worker创建一个单独的channel
	/*var channels [10]chan int
	  for i := 0; i < 10; i++ {
	      channels[i] = make(chan int)
	      // 创建一个goroutine来接受数据 => main goroutine <=channel=> 当前goroutine
	      go worker(i, channels[i])
	  }
	  for i := 0; i < 10; i++ {
	      // 向channel发送数据
	      channels[i] <- i
	  }*/

	for i := 0; i < 10; i++ {
		c := createWorker(i)
		c <- i
	}
	time.Sleep(1 * time.Second)
}

func BufferedChannelDemo() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	// 在这之前不需要接收者, 且不会发送deadlock
	// c <- 4 发送deadlock
	go worker(0, c)
	// 告知接收方数据发送完毕, 但是在main协程未结束时, 接收方仍然可以接受到默认值(例如int的默认值为0), 所以在接受channel
	// 数据时需要进行额外判断n, ok := <- c 或使用range判断
	close(c)
	time.Sleep(time.Millisecond)
}
