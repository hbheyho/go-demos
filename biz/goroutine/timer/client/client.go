package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 用net.Dial就可以简单地创建一个TCP连接
	con, err := net.Dial("tcp", "localhost:9000")
	defer con.Close()
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, con)
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Println(err)
		panic("copy error")
	}
}
