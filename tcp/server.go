package tcp

import (
	"fmt"
	"net"
	"time"
)

func Run() {
	l, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := l.Accept()
		go func(conn net.Conn) {
			buf := make([]byte, 1024)
			conn.Read(buf)
			fmt.Println(string(buf))

			conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
			time.Sleep(5 * time.Second)
			conn.Close()
		}(conn)
	}
}
