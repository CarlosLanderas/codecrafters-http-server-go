package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Request_From_Bytes(t *testing.T) {
	rawReq := "GET /index.html HTTP/1.1\r\nHost: localhost:4221\r\nUser-Agent: curl/7.64.1\r\nAccept: */*\r\n\r\n"

	req := RequestFromBytes([]byte(rawReq))

	host := req.Headers["Host"]
	userAgent := req.Headers["User-Agent"]

	assert.Equal(t, "GET", req.Method)
	assert.Equal(t, "/index.html", req.Path)
	assert.Equal(t, "HTTP/1.1", req.Protocol)
	assert.Equal(t, host, req.Headers["Host"])
	assert.Equal(t, userAgent, req.Headers["User-Agent"])
}

func Test_Request_Accept_Encoding(t *testing.T) {
	rawReq := "GET /index.html HTTP/1.1\r\nAccept-Encoding:invalid1, gzip, invalid2\r\n\r\n"
	req := RequestFromBytes([]byte(rawReq))

	assert.Equal(t, "gzip", req.AcceptEncoding()[0])
}
