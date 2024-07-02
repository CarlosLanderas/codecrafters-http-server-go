package http

import (
	"fmt"
	"regexp"
)

type Router struct {
	handlers map[string]RouteDefinition
}

type RouteDefinition struct {
	method  string
	handler RouteHandler
}

func (rd *RouteDefinition) Method(method string) {
	rd.method = method
}

type RouteHandler = func(*ResponseWriter, *HttpRequest) error

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]RouteDefinition),
	}
}

func (r *Router) Register(path string, method string, handler RouteHandler) {
	if _, ok := r.handlers[path]; !ok {
		r.handlers[r.HandlerKeyFromRequest(method, path)] = RouteDefinition{
			method:  method,
			handler: handler}
	}
}

func (r *Router) Get(req *HttpRequest) (RouteHandler, error) {

	for route := range r.handlers {
		match, err := regexp.Match(route, []byte("^"+r.HandlerKeyFromRequest(req.Method, req.Path)))
		if err != nil {
			return nil, err
		}

		if match {
			return r.handlers[route].handler, nil
		}
	}

	return nil, nil
}

func (r *Router) HandlerKeyFromRequest(method, path string) string {
	return fmt.Sprintf("%s_%s", method, path)
}
