package deuterium

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
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
	name, version, description, author string
	modules                            []*Module
	routes                             []*route
	middlewares                        []ContextHandler
	handlerIndex                       int
}

// Creates a new Deuterium Application with the
// specified modules.
func NewApp(modules []*Module) *app {
	return &app{
		modules: modules,
		routes:  []*route{},
	}
}

// Sets the name for the Deuterium Application
func (a *app) SetName(name string) {
	a.name = name
}

// Sets the verstion for the Deuterium Application
func (a *app) SetVersion(major, minor, build int) {
	builder := strings.Builder{}
	builder.WriteString("v")
	builder.WriteString(strconv.Itoa(major))
	builder.WriteString(".")
	builder.WriteString(strconv.Itoa(minor))
	builder.WriteString(".")
	builder.WriteString(strconv.Itoa(build))

	a.version = builder.String()
}

// Sets the description for the Deuterium Application
func (a *app) SetDescription(desc string) {
	a.description = desc
}

// Handles a GET request to the specified pattern and
// returns a pointer to a [route] instance that will
// be used to attach the endpoint handler.
func (a *app) Get(pattern string) *route {
	r := &route{
		method:  http.MethodGet,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

// Handles a POST request to the specified pattern and
// returns a pointer to a [route] instance that will
// be used to attach the endpoint handler.
func (a *app) Post(pattern string) *route {
	r := &route{
		method:  http.MethodPost,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

// Handles a PUT request to the specified pattern and
// returns a pointer to a [route] instance that will
// be used to attach the endpoint handler.
func (a *app) Put(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

// Handles a PATCH request to the specified pattern and
// returns a pointer to a [route] instance that will
// be used to attach the endpoint handler.
func (a *app) Patch(pattern string) *route {
	r := &route{
		method:  http.MethodPut,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

// Handles a DELETE request to the specified pattern and
// returns a pointer to a [route] instance that will
// be used to attach the endpoint handler.
func (a *app) Delete(pattern string) *route {
	r := &route{
		method:  http.MethodDelete,
		pattern: strings.TrimPrefix(pattern, "/"),
	}
	a.routes = append(a.routes, r)
	return r
}

// Takes a [ContextHandler] handler function which will
// be appended to the queue of middleware handlers for this application.
// To call the next handler, use [Context.Next()].
func (a *app) Use(handler ContextHandler) {
	a.middlewares = append(a.middlewares, handler)
}

func (a *app) ConfigureCors(cfg CorsConfig) {

}

func (a *app) next(ctx Context) {
	middlewareCount := len(a.middlewares)

	if middlewareCount == 0 {
		return
	}

	if a.handlerIndex < middlewareCount-1 {
		a.handlerIndex++
		a.middlewares[a.handlerIndex](ctx)
		return
	}

	return
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

		if len(a.middlewares) > 0 {
			a.middlewares[0](&context{
				req: &request{r, nil},
				res: &response{w},
				app: a,
			})

			if a.handlerIndex < len(a.middlewares)-1 {
				return
			}
		}

		var apiDocsEndpoints []string
		var match *route
		for i := 0; i < len(allRoutes); i++ {
			route := allRoutes[i]
			apiDocsEndpoints = append(apiDocsEndpoints, route.method+" "+route.pattern)
			matches := route.regexPattern.MatchString(r.URL.Path)
			if matches && route.method == r.Method {
				match = route
			}
		}

		if regexp.MustCompile(`api\-docs\/*`).MatchString(r.URL.Path) {
			bytes, _ := json.Marshal(struct {
				AppName   string   `json:"App Name"`
				Version   string   `json:"Version"`
				Endpoints []string `json:"Endpoints"`
			}{
				AppName:   a.name,
				Version:   a.version,
				Endpoints: apiDocsEndpoints,
			})

			w.Write(bytes)
			return
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
			app:   a,
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
