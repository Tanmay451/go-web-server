package tcp

import (
	"fmt"
	"net"
)

func Run() {
	l, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := l.Accept()

		buf := make([]byte, 1024)
		conn.Read(buf)
		fmt.Println(string(buf))

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

		conn.Close()
	}
}
