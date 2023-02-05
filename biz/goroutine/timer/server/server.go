package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Println(err)
		panic("listen error")
	}
	for {
		// listener对象的Accept方法会直接阻塞, 直到一个新的连接被创建, 然后会返回一个net.Conn对象来表示这个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			panic("accept error")
		}
		go handleCon(conn)
	}
}

// HandleCon 定时写入时间
func handleCon(con net.Conn) {
	defer con.Close()
	for {
		// 扩展: 输出流和输入流
		// 输入输出的方向是针对程序而言，向程序中读入数据，就是输入流；从程序中向外写出数据，就是输出流
		// 从磁盘、网络、键盘读到内存，就是输入流，用 Reader
		// 写到磁盘、网络、屏幕，都是输出流，用 Writer
		_, err := io.WriteString(con, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println(err)
			panic("write error")
			return
		}
		time.Sleep(1 * time.Second)
	}
}
