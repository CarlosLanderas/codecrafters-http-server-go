package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const Protocol = "HTTP/1.1"

func Test_Response(t *testing.T) {
	ok := OkResponse(Protocol, nil)
	nf := NotFoundResponse(Protocol, nil)

	assert.Equal(t, 200, ok.StatusCode)
	assert.Equal(t, 404, nf.StatusCode)
}

func Test_Response_Bytes(t *testing.T) {
	ok := OkResponse(Protocol, nil)

	resp := string(ok.Bytes())

	assert.Equal(t, resp, "HTTP/1.1 200 OK\r\n\r\n")
}

func Test_Response_Content(t *testing.T) {

	payload := "carlos"
	ok := OkResponse(Protocol, []byte(payload))
	ok.SetContentType("text/plain")

	resp := string(ok.Bytes())

	expectedResponse := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:6\r\n\r\ncarlos"

	assert.Equal(t, expectedResponse, resp)
}
