package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var newYork = flag.String("NewYork", "localhost:8010", "NewYork timezone ")
var tokyo = flag.String("Tokyo", "localhost:8020", "Tokyo timezone port")
var london = flag.String("London", "localhost:8030", "London timezone port")

func main() {
	flag.Parse()
	c := make(chan *string, 3)
	go func() {
		c <- newYork
		c <- tokyo
		c <- london
	}()
	for ch := range c {
		conn, err := net.Dial("tcp", *ch)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
