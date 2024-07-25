package http

import (
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Response_Writer(t *testing.T) {
	r, w := net.Pipe()

	defer r.Close()
	defer w.Close()

	writer := ResponseWriter{Conn: w}
	resp := NewResponse(200, "HTTP/1.1", []byte("lorem ipsum"))
	respBytes := resp.Bytes()

	go func() {
		writer.Write(resp)
	}()

	buff := make([]byte, len(respBytes))
	io.ReadFull(r, buff)

	assert.Equal(t, buff, respBytes)
}
