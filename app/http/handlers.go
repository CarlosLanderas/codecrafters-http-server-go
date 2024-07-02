package http

import (
	"net"
	"strings"
)

func RootHandler(conn net.Conn, req *HttpRequest) error {
	_, err := conn.Write(OkResponse(req.Protocol, nil).Bytes())
	return err
}

func NotFoundHandler(conn net.Conn, req *HttpRequest) error {
	_, err := conn.Write([]byte(NotFoundResponse(req.Protocol, nil).Bytes()))
	return err

}

func EchoHandler(conn net.Conn, req *HttpRequest) error {

	parts := strings.Split(req.Path, "/")

	content := parts[len(parts)-1]

	resp := OkResponse(req.Protocol, []byte(content))
	resp.SetContentType("text/plain")

	_, err := conn.Write(resp.Bytes())

	return err
}

func UserAgentHandler(conn net.Conn, req *HttpRequest) error {
	userAgent := req.Headers["User-Agent"]
	resp := OkResponse(req.Protocol, []byte(userAgent))
	resp.SetContentType("text/plain")

	_, err := conn.Write(resp.Bytes())
	return err
}
