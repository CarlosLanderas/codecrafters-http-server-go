package http

import (
	"bytes"
	"io"
	"slices"
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

	request := &HttpRequest{}

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

	request.Method = actionParts[0]
	request.Path = actionParts[1]
	request.Protocol = actionParts[2]
	request.Headers = headers

	body := reqParts[headerEnd+1]
	content := body[:request.ContentLength()]

	request.Body = bytes.NewReader([]byte(content))

	return request
}

func (hr *HttpRequest) ContentLength() int {
	val, ok := hr.Headers["Content-Length"]

	if ok {
		length, err := strconv.Atoi(val)
		if err != nil {
			return 0
		}
		return length
	}

	return 0
}

func (hr *HttpRequest) ValidEncoding() bool {
	value := hr.AcceptEncoding()

	return slices.ContainsFunc(SupportedEncodings, func(enc string) bool {
		return slices.Contains(value, enc)
	})
}

func (hr *HttpRequest) AcceptEncoding() []string {
	val, ok := hr.Headers["Accept-Encoding"]

	encodings := make([]string, 0)

	if !ok {
		return encodings
	}

	if strings.Contains(val, ",") {
		encodings = strings.Split(val, ",")
	} else {
		encodings = []string{val}
	}

	for i, e := range encodings {
		encodings[i] = strings.TrimSpace(e)
	}

	validEnc := make([]string, 0)

	for _, e := range encodings {
		if slices.Contains(SupportedEncodings, e) {
			validEnc = append(validEnc, e)
		}
	}

	return validEnc
}
