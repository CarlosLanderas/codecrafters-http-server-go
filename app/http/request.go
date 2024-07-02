package http

import "strings"

type HttpRequest struct {
	Protocol string
	Method   string
	Path     string
	Headers  map[string]string
}

func RequestFromBytes(buff []byte) *HttpRequest {

	req := string(buff)
	reqParts := strings.Split(req, "\r\n")
	actionParts := strings.Split(reqParts[0], " ")

	// Parse headers
	headers := make(map[string]string)

	for i := 1; i < len(reqParts)-1; i++ {
		if reqParts[i] == "" {
			continue
		}

		headerParts := strings.SplitN(reqParts[i], ":", 2)
		headers[headerParts[0]] = strings.TrimSpace(headerParts[1])
	}

	return &HttpRequest{
		Method:   actionParts[0],
		Path:     actionParts[1],
		Protocol: actionParts[2],
		Headers:  headers,
	}
}
