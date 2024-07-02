package server

import (
	"fmt"
	"net"
)

type Server struct {
	port     int
	listener net.Listener
}

func New(port int) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	var err error

	fmt.Println("Starting server on port:", s.port)

	s.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", s.port))

	if err != nil {
		return fmt.Errorf("error binding port: %v", err)
	}

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn)
	}
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 200 message
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
