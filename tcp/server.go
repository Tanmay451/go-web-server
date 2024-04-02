package tcp

import "net"

func Run() {
	l, _ := net.Listen("tcp", ":8080")
	conn, _ := l.Accept()
	conn.Close()
}
