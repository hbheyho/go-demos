package main

import (
	"fmt"
	"time"
)

/* go并发编程 - goroutine
1. 通过go关键字来完成协程(Coroutine)的创建, go语言原生支持协程;
2. main协程和goroutine并发执行, 当main退出, goroutine也停止运行;
3. 协程的运行顺序不定, 通过goroutine调度器进行调度;
4. 通过runtime.Gosched手动进行让出当前协程的执行时间;
5. 子程序是协程的一个特例, 协程是一个比子程序更加宽泛的概念;
6. go调度器会在合适的点进行切换, 这也和传统的协程存在区别, 传统协程需要显示定义切换点
     => go调度器切换点：I/O, select, channel, 等待锁, 函数调用(有时), runtime.Gosched
    虽然提供了上述切换点, 但是总体上来说, goroutine还是一个非抢占式多任务处理.
【扩展】：1. 协程是 一个轻量级的"线程", 非抢占式多任务处理, 由协程主动交出控制权, 编译器/解释器/虚拟机层面的多任务,
           而不是操作系统级别的多任务, 多个协程可以在一个或多个线程上运行.
        2. 在go 1.14之后引入了基于系统信号的异步抢占调度, 死循环的goroutine可以被抢占调度
           参考文档：https://mp.weixin.qq.com/s/PVxdtvSXgNpiD65TUo-TCg
*/
func main() {
	// goroutine1()
	goroutine2()
	time.Sleep(10 * time.Second)
	fmt.Println(a)
}

func goroutine1() {
	for i := 0; i < 10; i++ {
		// 协程的运行顺序不定, 可能 goroutine 1先比 goroutine 0运行
		go func(i int) {
			for {
				// Printf作为一个IO操作, 所以就算是一个for语句死循坏也会出现协程切换
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
}

var a [10]int

func goroutine2() {
	for i := 0; i < 10; i++ {
		// 协程的运行顺序不定, 可能 goroutine 1先比 goroutine 0运行
		// 协程是非抢占式多任务处理, 那么当一个协程得到机会运行, 且运行的代码为死循坏(类似下文的for循环), 那么会造成其他协程
		// 无法执行而导致宕机
		// 但是这种情况只会出现在go 1.14之前, 在go 1.14之后引入了基于系统信号的异步抢占调度, 下面的死循环的goroutine可以被抢占调度
		// 参考文档：https://mp.weixin.qq.com/s/PVxdtvSXgNpiD65TUo-TCg
		go func(i int) {
			a[i]++
		}(i)
		// index out of range [10] with length 10. 当i = 10跳出循环时, goroutine仍然执行, 其又在引用外部的i, 所有
		// 出现index out of range错误
		//go func() {
		//    for {
		//        a[i]++
		//    }
		//}()
	}
}
