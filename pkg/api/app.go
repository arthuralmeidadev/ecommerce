package api

import (
	"fmt"
	"net/http"
	"strings"
)

func mapRoutes(modules []*Module) http.Handler {
	logger := GetLogger()
	var allRoutes []*route
	var deps []any
	for _, m := range modules {
		routes, providers := m.Register()
		for _, route := range routes {
			if !route.registered {
				logger.Warn(fmt.Sprintf("route \"%s\" needs to be registered!", route.pattern))
				continue
			}
			allRoutes = append(allRoutes, route)
		}

		deps = append(deps, providers...)
	}
	c := GetContainer()
	if err := c.BulkInject(deps); err != nil {
		logger.Warn(fmt.Sprintf("Could not inject dependencies: \n%v", err))
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

		var matchedRoute *route
		for _, route := range allRoutes {
			if route.regexPattern.MatchString(r.URL.Path) && route.method == r.Method {
				matchedRoute = route
			}
		}

		if matchedRoute == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		reqPathParams := make(map[string]string)
		incomingSegments := strings.Split(r.URL.Path, "/")
		for i, segment := range strings.Split(matchedRoute.pattern, "/") {
			if strings.HasPrefix(segment, ":") && len(segment) > 1 {
				reqPathParams[segment[1:]] = incomingSegments[i+1]
			}
		}

		matchedRoute.handler(
			&context{
				req: &request{r, reqPathParams},
				res: &response{w},
			},
		)
	})
}

type app struct {
	handler http.Handler
	modules []*Module
}

func NewApp(modules []*Module) *app {
	return &app{
		handler: mapRoutes(modules),
		modules: modules,
	}
}

func (a *app) Listen(address string, port int) {
	logger := GetLogger()
	ctlChan := make(chan struct{})
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), a.handler)
		if err != nil {
			ctlChan <- struct{}{}
			logger.Fatal(fmt.Sprintf("Coudn't start app: %v", err))
		}
	}()
    launch(fmt.Sprintf("APP LAUNCHED: Listening on port %d.", port))
	<-ctlChan
}
