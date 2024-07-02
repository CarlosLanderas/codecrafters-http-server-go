package http

import (
	"net"
	"regexp"
)

type Router struct {
	handlers map[string]RouteHandler
}

const (
	LiteralRoute = iota
	RegexRouter
)

type RouteHandler = func(net.Conn, *HttpRequest) error

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]RouteHandler),
	}
}

func (r *Router) Register(path string, handler RouteHandler) {
	if _, ok := r.handlers[path]; !ok {
		r.handlers[path] = handler
	}
}

func (r *Router) Get(path string) (RouteHandler, error) {

	for route := range r.handlers {
		match, err := regexp.Match(route, []byte(path))
		if err != nil {
			return nil, err
		}

		if match {
			return r.handlers[route], nil
		}
	}

	return nil, nil
}
