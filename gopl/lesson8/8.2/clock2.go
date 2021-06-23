package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 非并发
func main() {
	// 监听一个端口
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 接收请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 关于请求的处理
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	// 定时输出时间
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
