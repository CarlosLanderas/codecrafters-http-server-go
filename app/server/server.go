package server

import (
	"fmt"
	"log"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
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

func handleConnection(conn net.Conn) error {
	defer conn.Close()

	// Read request
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)

	if err != nil {
		log.Fatalf("error reading request: %v", err)
	}

	req := http.RequestFromBytes(buff)

	if req.Path == "/" {
		conn.Write(http.OkResponse(req.Protocol).Bytes())
	} else {
		conn.Write([]byte(http.NotFoundResponse(req.Protocol).Bytes()))
	}

	fmt.Printf("%+v", req)

	return nil
}
