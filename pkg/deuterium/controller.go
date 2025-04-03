package deuterium

import (
	"net/http"
	"strings"
)

type Controller interface {
	Get(pattern string) *route
	Post(pattern string) *route
	Put(pattern string) *route
	Patch(pattern string) *route
	Delete(pattern string) *route
	Use(handler ContextHandler)
	register() ([]*route, []ContextHandler)
}

type ControllerFactory interface {
	Make() Controller
}

type controller struct {
	baseRoute    string
	routes       []*route
	middlewares  []ContextHandler
	handlerIndex int
}

func NewController(baseRoute string) Controller {
	return &controller{
		baseRoute: strings.TrimSuffix(strings.TrimPrefix(baseRoute, "/"), "/"),
		routes:    make([]*route, 0),
	}
}

func (c *controller) Get(pattern string) *route {
	r := &route{
		method:  http.MethodGet,
		pattern: c.baseRoute + "/" + strings.TrimPrefix(pattern, "/"),
	}
	c.routes = append(c.routes, r)
	return r
}

func (c *controller) Post(pattern string) *route {
	r := &route{
		method:  http.MethodPost,
		pattern: c.baseRoute + "/" + strings.TrimPrefix(pattern, "/"),
	}
	c.routes = append(c.routes, r)
	return r
}

func (c *controller) Put(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: c.baseRoute + "/" + strings.TrimPrefix(pattern, "/"),
	}
	c.routes = append(c.routes, r)
	return r
}

func (c *controller) Patch(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: c.baseRoute + "/" + strings.TrimPrefix(pattern, "/"),
	}
	c.routes = append(c.routes, r)
	return r
}

func (c *controller) Delete(pattern string) *route {
	r := &route{
		method:  http.MethodDelete,
		pattern: c.baseRoute + "/" + strings.TrimPrefix(pattern, "/"),
	}
	c.routes = append(c.routes, r)
	return r
}

func (c *controller) Use(handler ContextHandler) {
	c.middlewares = append(c.middlewares, handler)
}

func (c *controller) register() ([]*route, []ContextHandler) {
	return c.routes, c.middlewares
}
