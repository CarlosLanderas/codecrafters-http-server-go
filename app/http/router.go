package http

import (
	"regexp"
)

type Router struct {
	handlers map[string]RouteHandler
}

type RouteHandler = func(*ResponseWriter, *HttpRequest) error

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
