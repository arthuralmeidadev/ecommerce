package api

import (
	"log"
	"net/http"
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
	log.Printf("\x1b[92mRegistered route: %s %s\x1b[0m", r.method, r.pattern)
}

type router struct {
	routes []*route
}

func NewRouter() *router {
	return &router{
		routes: make([]*route, 0),
	}
}

func (r *router) Get(pattern string) *route {
	rt := &route{
		method:  http.MethodGet,
		pattern: pattern,
	}
	r.routes = append(r.routes, rt)
	return rt
}

func (r *router) Post(pattern string) *route {
	rt := &route{
		method:  http.MethodPost,
		pattern: pattern,
	}
	r.routes = append(r.routes, rt)
	return rt
}

func (r *router) Put(pattern string) *route {
	rt := &route{
		method:  http.MethodPut,
		pattern: pattern,
	}
	r.routes = append(r.routes, rt)
	return rt
}

func (r *router) Patch(pattern string) *route {
	rt := &route{
		method:  http.MethodPut,
		pattern: pattern,
	}
	r.routes = append(r.routes, rt)
	return rt
}

func (r *router) Delete(pattern string) *route {
	rt := &route{
		method:  http.MethodDelete,
		pattern: pattern,
	}
	r.routes = append(r.routes, rt)
	return rt
}

func (rt *router) Register() http.Handler {
	for _, route := range rt.routes {
		if !route.registered {
			log.Printf("\x1b[93mWARNING: route \"%s\" needs to be registered!\x1b[0m", route.pattern)
		}
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"OPTIONS, GET, POST, PUT, PATCH, PUT, DELETE",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With",
		)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		var matchedRoute *route
		for _, route := range rt.routes {
			if route.regexPattern.MatchString(r.URL.Path) && route.method == r.Method {
				matchedRoute = route
			}
		}

		if matchedRoute == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		reqPathParams := make(map[string]string, 0)
		incomingSegments := strings.Split(r.URL.Path, "/")
		for i, segment := range strings.Split(matchedRoute.pattern, "/") {
			if strings.HasPrefix(segment, ":") && len(segment) > 1 {
				reqPathParams[segment[1:]] = incomingSegments[i]
			}
		}

		matchedRoute.handler(
			&context{
				req: &request{r, reqPathParams},
				res: &response{w},
			},
		)
	})

	log.Println("\x1b[92mRouter successfully registered!\x1b[0m")
	return handler
}
