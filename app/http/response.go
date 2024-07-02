package http

import (
	"fmt"
	"net/http"
)

func NewResponse(statusCode int, protocol string) *HttpResponse {
	return &HttpResponse{
		Protocol:   protocol,
		StatusCode: statusCode,
	}
}

func OkResponse(protocol string, payload []byte) *HttpResponse {
	return &HttpResponse{
		Protocol:   protocol,
		Payload:    payload,
		StatusCode: http.StatusOK,
	}
}

func NotFoundResponse(protocol string, payload []byte) *HttpResponse {
	return &HttpResponse{
		Protocol:   protocol,
		Payload:    payload,
		StatusCode: http.StatusNotFound,
	}
}

type HttpResponse struct {
	Protocol    string
	StatusCode  int
	Payload     []byte
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
	if r.Payload != nil {
		return len(r.Payload)
	}

	return 0
}

func (r *HttpResponse) String() string {
	return fmt.Sprintf("%s %d %s\r\nContent-Type: %s\r\nContent-Length:%d\r\n\r\n%s",
		r.Protocol,
		r.StatusCode,
		http.StatusText(r.StatusCode),
		r.ContentType,
		r.Length(),
		string(r.Payload))
}

func (r *HttpResponse) SetContentType(contentType string) {
	r.ContentType = contentType
}