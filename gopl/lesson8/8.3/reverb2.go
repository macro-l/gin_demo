package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

func echo(c net.Conn, shout string, delay time.Duration) {
	// 炒鸡大声
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	// 声音肥来了
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	// 小声逼逼
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	// 再额外配个环绕音
	fmt.Fprintln(c, "\t....")
}

func handleConn(c net.Conn) {
	defer c.Close()

	// 嗅探输入
	input := bufio.NewScanner(c)

	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

}
