package http

import "strings"

type HttpRequest struct {
	Protocol string
	Method   string
	Path     string
}

func RequestFromBytes(buff []byte) *HttpRequest {

	req := string(buff)
	reqParts := strings.Split(req, "\r\n")
	actionParts := strings.Split(reqParts[0], " ")

	return &HttpRequest{
		Method:   actionParts[0],
		Path:     actionParts[1],
		Protocol: actionParts[2],
	}
}
