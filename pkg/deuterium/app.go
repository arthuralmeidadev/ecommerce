package deuterium

import (
	"fmt"
	"net/http"
	"strings"
)

func appendRoutes(allRoutes *[]*route, routes []*route, logger *Logger) {
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		if !route.registered {
			msg := fmt.Sprintf("route \"%s\" needs to be registered!", route.pattern)
			logger.Warn(msg)
			continue
		}
		*allRoutes = append(*allRoutes, route)
	}
}

type app struct {
	modules []*Module
	routes  []*route
}

func NewApp(modules []*Module) *app {
	return &app{
		modules: modules,
		routes:  []*route{},
	}
}

func (a *app) Get(pattern string) *route {
	r := &route{
		method:  http.MethodGet,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

func (a *app) Post(pattern string) *route {
	r := &route{
		method:  http.MethodPost,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

func (a *app) Put(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

func (a *app) Patch(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

func (a *app) Delete(pattern string) *route {
	r := &route{
		method:  http.MethodDelete,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

func (a *app) register() http.Handler {
	logger := GetLogger()
	allRoutes := a.routes
	var deps []any
	for i := 0; i < len(a.modules); i++ {
		routes, providers := a.modules[i].Register()
		appendRoutes(&allRoutes, routes, logger)
		deps = append(deps, providers...)
	}

	c := GetContainer()
	if err := c.BulkInject(deps); err != nil {
		logger.Fatal(fmt.Sprintf("Could not inject dependencies:\n\t%v", err))
	}
	logger.Success("All required dependencies loaded!")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		var match *route
		for i := 0; i < len(allRoutes); i++ {
			route := allRoutes[i]
			matches := route.regexPattern.MatchString(r.URL.Path)
			if matches && route.method == r.Method {
				match = route
			}
		}

		if match == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		reqPathParams := make(map[string]string)
		incomingSegments := strings.Split(r.URL.Path, "/")
		for i, segment := range strings.Split(match.pattern, "/") {
			if strings.HasPrefix(segment, ":") && len(segment) > 1 {
				reqPathParams[segment[1:]] = incomingSegments[i+1]
			}
		}

		ctx := &context{
			req:   &request{r, reqPathParams},
			res:   &response{w},
			route: match,
		}

		if len(match.middlewares) > 0 {
			match.middlewares[0](ctx)
			return
		}

		match.handler(ctx)
	})
}

func (a *app) Listen(addr string, port int) {
	logger := GetLogger()
	ctlChan := make(chan struct{})
	handler := a.register()
	go func() {
		addr = fmt.Sprintf("%s:%d", addr, port)
		err := http.ListenAndServe(addr, handler)
		if err != nil {
			ctlChan <- struct{}{}
			logger.Fatal(fmt.Sprintf("Coudn't start app: %v", err))
		}
	}()
	launch(fmt.Sprintf("App listening on port %d.", port))
	<-ctlChan
}
