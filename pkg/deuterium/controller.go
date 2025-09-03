package deuterium

import (
	"net/http"
	"strings"
)

type Controller interface {
	// Handles a GET request to the specified pattern and
	// returns a pointer to a [route] instance that will
	// be used to attach the endpoint handler.
	Get(pattern string) *route

	// Handles a POST request to the specified pattern and
	// returns a pointer to a [route] instance that will
	// be used to attach the endpoint handler.
	Post(pattern string) *route

	// Handles a PUT request to the specified pattern and
	// returns a pointer to a [route] instance that will
	// be used to attach the endpoint handler.
	Put(pattern string) *route

	// Handles a PATCH request to the specified pattern and
	// returns a pointer to a [route] instance that will
	// be used to attach the endpoint handler.
	Patch(pattern string) *route

	// Handles a DELETE request to the specified pattern and
	// returns a pointer to a [route] instance that will
	// be used to attach the endpoint handler.
	Delete(pattern string) *route

	// Takes a [ContextHandler] handler function which will
	// be appended to the queue of middleware handlers for this route.
	// To call the next handler, use [Context.Next()].
	Use(handler ContextHandler)
	register() ([]*route, []ContextHandler)
}

type controller struct {
	baseRoute    string
	routes       []*route
	middlewares  []ContextHandler
	handlerIndex int
}

// Creates a new struct that implements [Controller]
// associated with the base route.
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
