package main

import (
	"fmt"
	"math/rand"
	"time"
)

// select:
// 1. select 能够避免因为发送或者接收channel导致的阻塞，尤其是当channel没有准备好写或者读时
// 2. 通过select来实现选择性数据收发, 同时监听多个channel, 当某个channel数据到达时就进行数据处理
// 3. select会有一个default来设置当其它的操作都不能够马上被处理时程序需要执行哪些逻辑
// 4. 每一个case代表一个通信操作（在某个channel上进行发送或者接收），并且会包含一些语句组成的一个语句块
// 5. 如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会
// 6. channel的零值是nil.  在select中可以使用nil channel, 对一个nil的channel发送和接收操作会永远阻塞，在select语句中操作nil的channel永远都不会被select到。
// 7. 通过time.After(time)实现定时器的作用, time时间之后返回往<-chan Time送入一个值
// 8. 通过time.Tick(time)实现间隔定时效果, 每隔time时间返回往<-chan Time送入一个值
// 9. case 条件中不仅仅可以是channel接收操作, 也可以是发送操作
func main() {
	// 生产数据到c1, c2
	var c1, c2 = generate(), generate()
	// 创建一个消费者, 从c1, c2消费数据
	var worker = createWorker(0)
	// 存储c1, c2接收到的数据, 防止出现生产者和消费者速度不一致问题
	// 相当于缓存队列
	var values []int

	// tm类型为： <-chan Time => 实现类似定时器效果
	// 10s中之后定时退出
	tm := time.After(10 * time.Second)
	// tick类型为： <-chan Time
	// 每隔1s之后就返回
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		// 缓存队列有数据了, 开始进行消费
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:] // 从缓存队列删除values[0]
		case <-time.After(time.Millisecond * 800): // 定义一个定时器：当两个channel之间的数据间隔超过800ms, 就显示超时
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("buffer len = ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}

// worker 数据消费者
func worker(id int, c chan int) {
	for x := range c {
		fmt.Printf("worker %d received %d \n", id, x)
	}
}

// createWorker 创建一个消费者等待消费, 并将channel透出
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// generate 数据生产者, 生产数据到chan out, 并透出给调用方
func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for true {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
