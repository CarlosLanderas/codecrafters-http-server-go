package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Router(t *testing.T) {
	router := NewRouter()

	executed := false

	router.Register("/echo/.*", "GET", func(w *ResponseWriter, hr *HttpRequest) error {
		executed = true
		return nil
	})

	req := &HttpRequest{
		Method: "GET",
		Path:   "/echo/carlos",
	}

	route, err := router.Get(req)

	if err != nil {
		panic(err)
	}

	route(nil, nil)

	assert.Equal(t, true, executed)
}

func Test_Route_Not_Found(t *testing.T) {
	router := NewRouter()

	router.Register("/echo/.*", "GET", func(w *ResponseWriter, hr *HttpRequest) error {
		return nil
	})

	req := &HttpRequest{
		Method: "POST",
		Path:   "/echo/carlos",
	}

	route, _ := router.Get(req)

	assert.Nil(t, route)
}
