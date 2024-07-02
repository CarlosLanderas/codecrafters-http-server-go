package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Protocol = "HTTP/1.1"

func Test_Response(t *testing.T) {
	ok := OkResponse(Protocol)
	nf := NotFoundResponse(Protocol)

	assert.Equal(t, 200, ok.StatusCode)
	assert.Equal(t, 404, nf.StatusCode)
}

func Test_Response_Bytes(t *testing.T) {
	ok := OkResponse(Protocol)

	resp := string(ok.Bytes())

	assert.Equal(t, resp, "HTTP/1.1 200 OK\r\n\r\n")
}
