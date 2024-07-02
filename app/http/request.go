package http

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type HttpRequest struct {
	Protocol string
	Method   string
	Path     string
	Body     io.Reader
	Headers  map[string]string
}

func RequestFromBytes(buff []byte) *HttpRequest {

	req := string(buff)
	reqParts := strings.Split(req, "\r\n")
	actionParts := strings.Split(reqParts[0], " ")

	// Parse headers
	headers := make(map[string]string)

	headerEnd := 0

	for i := 1; i < len(reqParts)-1; i++ {
		if reqParts[i] == "" {
			headerEnd = i
			continue

		}

		headerParts := strings.SplitN(reqParts[i], ":", 2)
		headers[headerParts[0]] = strings.TrimSpace(headerParts[1])
	}

	// Read the body

	contentLength := 0

	if length, ok := headers["Content-Length"]; ok {
		length, err := strconv.Atoi(length)
		if err == nil {
			contentLength = length
		}
	}

	body := reqParts[headerEnd+1]

	content := body[:contentLength]

	return &HttpRequest{
		Method:   actionParts[0],
		Path:     actionParts[1],
		Protocol: actionParts[2],
		Body:     bytes.NewReader([]byte(content)),
		Headers:  headers,
	}
}
