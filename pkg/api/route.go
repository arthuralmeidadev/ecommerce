package api

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
	handler      ContextHandler
}

func (r *route) Register(ctxHandler ContextHandler) {
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
	r.handler = ctxHandler
	r.registered = true
	r.regexPattern = *regexp.MustCompile(builder.String())
	logger := GetLogger()
	logger.Success(fmt.Sprintf("Registered route: %s %s", r.method, r.pattern))
}
