package tcp

import (
	"fmt"
	"net"
	"time"
)

type Server struct {
	addr string
	ln   net.Listener
}

func NewServer(add string) *Server {
	return &Server{
		addr: add,
	}
}

func (s *Server) Start() error {
	ln, _ := net.Listen("tcp", s.addr)
	s.ln = ln
	s.accept()
	return nil
}

func (s *Server) accept() error {
	for {
		conn, _ := s.ln.Accept()
		go s.read(conn)
	}
}

func (s *Server) read(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(15 * time.Second))
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	buf := make([]byte, 1024)
	conn.Read(buf)
	fmt.Println(string(buf))

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	time.Sleep(5 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	time.Sleep(5 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	time.Sleep(5 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))
	conn.Close()
}

func Run() {
	server := NewServer(":3000")
	server.Start()
}
