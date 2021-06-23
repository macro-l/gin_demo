/*
   练习 8.1
   实现 clockwall
*/
package exercise8_1

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "8000", "different timezone port")

func main() {
	//fmt.Printf("%+v", *port)
	//os.Exit(0)

	// 标志位解析
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
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
		go handleConn(conn, *port)
	}
}

func handleConn(c net.Conn, port string) {
	defer c.Close()

	var l *time.Location
	switch port {
	case "8010":
		l, _ = time.LoadLocation("US/Eastern") //根据port更改时区
	case "8020":
		l, _ = time.LoadLocation("Asia/Tokyo")
	case "8030":
		l, _ = time.LoadLocation("Europe/London")
	}
	// 定时输出时间
	for {
		_, err := io.WriteString(c, l.String()+":"+time.Now().In(l).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
