package deuterium

import (
	"fmt"
	"regexp"
	"strings"
)

type ContextHandler func(ctx Context)

type route struct {
	method       string
	pattern      string
	regexPattern regexp.Regexp
	registered   bool
	middlewares  []ContextHandler
	handlerIndex int
	handler      ContextHandler
}

func (r *route) next(ctx Context) {
	middlewareCount := len(r.middlewares)

	if middlewareCount == 0 || r.handlerIndex > middlewareCount-1 {
		return
	}

	if r.handlerIndex < middlewareCount-1 {
		r.handlerIndex++
		r.middlewares[r.handlerIndex](ctx)
		return
	}

	r.handlerIndex++
	r.handler(ctx)
}

// Takes a [ContextHandler] handler function which will
// be appended to the queue of middleware handlers for this route.
// To call the next handler, use [Context.Next()].
func (r *route) Use(handler ContextHandler) *route {
	r.middlewares = append(r.middlewares, handler)
	return r
}

// Takes a [ContextHandler] handler function as the main
// endpoint handler. This method needs to be called, otherwise
// the endpoint will be listed as unregistered.
func (r *route) Register(handler ContextHandler) {
	segments := strings.Split(r.pattern, "/")
	var builder strings.Builder
	builder.WriteString(`^`)

	for _, segment := range segments {
		if strings.HasPrefix(segment, ":") && len(segment) > 1 {
			builder.WriteString(`\/[^\/]+`)
		} else if len(segment) > 0 {
			builder.WriteString(`\/`)
			builder.WriteString(segment)
		}
	}

	builder.WriteString(`\/?`)
	r.handler = handler
	r.registered = true
	r.regexPattern = *regexp.MustCompile(builder.String())
	logger := GetLogger()
	logger.Success(fmt.Sprintf("Registered route: %s %s", r.method, r.pattern))
}
