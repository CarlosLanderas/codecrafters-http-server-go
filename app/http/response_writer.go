package http

import (
	"bytes"
	"compress/gzip"
	"net"
)

type ResponseWriter struct {
	Conn     net.Conn
	Encoding string
	response *HttpResponse
}

func (rw *ResponseWriter) Write(response *HttpResponse) (int, error) {
	rw.response = response

	if rw.Encoding != "" {
		// Compress
		response := rw.response
		response.Headers["Content-Encoding"] = rw.Encoding

		if rw.Encoding == "gzip" {
			response.SetPayload(gzipBody(response.payload))
		}

	}

	return rw.flush()
}

func (rw *ResponseWriter) flush() (int, error) {
	return rw.Conn.Write(rw.response.Bytes())
}

func gzipBody(body []byte) []byte {
	var buff bytes.Buffer
	gzip := gzip.NewWriter(&buff)
	gzip.Write(body)
	gzip.Close()

	return buff.Bytes()
}
