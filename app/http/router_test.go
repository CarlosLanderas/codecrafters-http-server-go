package http

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Router(t *testing.T) {
	router := NewRouter()

	executed := false

	router.Register("/echo/.*", func(c net.Conn, hr *HttpRequest) error {
		executed = true
		return nil
	})

	route, err := router.Get("/echo/carlos")

	if err != nil {
		panic(err)
	}

	route(nil, nil)

	assert.Equal(t, true, executed)
}
