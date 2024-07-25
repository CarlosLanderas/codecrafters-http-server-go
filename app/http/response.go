package http

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

func NewResponse(statusCode int, protocol string, payload []byte) *HttpResponse {
	r := &HttpResponse{
		Protocol:   protocol,
		StatusCode: statusCode,
		payload:    payload,
		Headers:    make(map[string]string),
	}

	r.SetPayload(payload)

	return r
}

func OkResponse(protocol string, payload []byte) *HttpResponse {
	return NewResponse(http.StatusOK, protocol, payload)
}

func CreatedResponse(protocol string, payload []byte) *HttpResponse {
	return NewResponse(http.StatusCreated, protocol, payload)
}

func NotFoundResponse(protocol string, payload []byte) *HttpResponse {
	return NewResponse(http.StatusNotFound, protocol, payload)
}

type HttpResponse struct {
	Protocol    string
	StatusCode  int
	payload     []byte
	ContentType string
	Headers     map[string]string
}

func (r *HttpResponse) AddHeader(name, value string) {
	r.Headers[name] = value
}

func (r *HttpResponse) Bytes() []byte {
	return []byte(r.String())
}

func (r *HttpResponse) Length() int {
	if r.payload != nil {
		return len(r.payload)
	}

	return 0
}

func (r *HttpResponse) SetPayload(bytes []byte) {
	r.payload = bytes

	if bytes != nil {
		r.SetContentLength(len(bytes))
	} else {
		r.SetContentLength(0)
	}
}

func (r *HttpResponse) String() string {
	return fmt.Sprintf("%s %d %s\r\n%s\r\n%s",
		r.Protocol,
		r.StatusCode,
		http.StatusText(r.StatusCode),
		r.renderHeaders(),
		string(r.payload))
}

func (r *HttpResponse) SetContentType(contentType string) {
	r.Headers["Content-Type"] = contentType
}

func (r *HttpResponse) SetContentLength(contentLength int) {
	length := strconv.Itoa(contentLength)
	r.Headers["Content-Length"] = length
}

func (r *HttpResponse) SetContentEncoding(encoding string) {
	r.Headers["Content-Encoding"] = encoding
}

func (r *HttpResponse) renderHeaders() string {
	var buff bytes.Buffer

	for k, v := range r.Headers {
		headerRaw := fmt.Sprintf("%s: %s\r\n", k, v)
		buff.WriteString(headerRaw)
	}

	return buff.String()
}
