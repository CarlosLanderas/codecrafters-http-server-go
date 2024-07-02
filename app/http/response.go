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

func OkResponse(protocol string) *HttpResponse {
	return &HttpResponse{
		Protocol:   protocol,
		StatusCode: http.StatusOK,
	}
}

func NotFoundResponse(protocol string) *HttpResponse {
	return &HttpResponse{
		Protocol:   protocol,
		StatusCode: http.StatusNotFound,
	}
}

type HttpResponse struct {
	Protocol   string
	StatusCode int
	Headers    map[string]string
}

func (r *HttpResponse) AddHeader(name, value string) {
	r.Headers[name] = value
}

func (r *HttpResponse) Bytes() []byte {
	return []byte(r.String())
}

func (r *HttpResponse) String() string {
	return fmt.Sprintf("%s %d %s\r\n\r\n", r.Protocol, r.StatusCode, http.StatusText(r.StatusCode))
}
