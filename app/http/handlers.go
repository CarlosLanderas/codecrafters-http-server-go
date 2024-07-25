package http

import (
	"io"
	"os"
	"path"
	"strings"
)

func RootHandler(w *ResponseWriter, req *HttpRequest) error {
	_, err := w.Write(OkResponse(req.Protocol, nil))
	return err
}

func NotFoundHandler(w *ResponseWriter, req *HttpRequest) error {
	_, err := w.Write(NotFoundResponse(req.Protocol, nil))
	return err
}

func EchoHandler(w *ResponseWriter, req *HttpRequest) error {
	parts := strings.Split(req.Path, "/")
	content := parts[len(parts)-1]
	resp := OkResponse(req.Protocol, []byte(content))
	resp.SetContentType("text/plain")

	_, err := w.Write(resp)

	return err
}

func UserAgentHandler(w *ResponseWriter, req *HttpRequest) error {
	userAgent := req.Headers["User-Agent"]
	resp := OkResponse(req.Protocol, []byte(userAgent))
	resp.SetContentType("text/plain")

	_, err := w.Write(resp)
	return err
}

func FileHandler(filePath string) func(w *ResponseWriter, req *HttpRequest) error {
	return func(w *ResponseWriter, req *HttpRequest) error {
		fileName := pathFileName(req.Path)
		path := path.Join(filePath, fileName)

		var resp *HttpResponse

		if _, err := os.Stat(path); os.IsNotExist(err) {
			resp = NotFoundResponse(req.Protocol, nil)
		} else {
			data, _ := os.ReadFile(path)
			resp = OkResponse(req.Protocol, data)
			resp.SetContentType("application/octet-stream")
		}

		_, err := w.Write(resp)
		return err
	}
}

func FilePostHandler(filePath string) func(w *ResponseWriter, req *HttpRequest) error {
	return func(w *ResponseWriter, req *HttpRequest) error {
		fileName := pathFileName(req.Path)
		path := path.Join(filePath, fileName)
		body, err := io.ReadAll(req.Body)

		if err != nil {
			return err
		}

		os.WriteFile(path, body, os.ModePerm)
		w.Write(CreatedResponse(req.Protocol, nil))

		return nil
	}
}

func pathFileName(path string) string {
	fileStart := strings.LastIndex(path, "/")
	fileName := path[fileStart+1:]

	return fileName
}
