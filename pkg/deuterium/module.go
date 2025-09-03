package deuterium

import "fmt"

type ModuleFactory interface {
	Make() *Module
}

type Module struct {
	Name       string
	Imports    []*Module
	Controller Controller
	Providers  []any
}

func (m *Module) Register() ([]*route, []any) {
	logger := GetLogger()
	var routes []*route
	var middlewares []ContextHandler
	if m.Controller != nil {
		routes, middlewares = m.Controller.register()
		if m.Name == "" {
			m.Name = "\"Unnamed\""
		}
		logger.Success(fmt.Sprintf("%s controller successfully registered", m.Name))
	}

	for i := 0; i < len(routes); i++ {
		var routeMiddlewares []ContextHandler
		routeMiddlewares = append(routeMiddlewares, middlewares...)
		routeMiddlewares = append(routeMiddlewares, routes[i].middlewares...)
		routes[i].middlewares = routeMiddlewares
	}

	providers := m.Providers
	for i := 0; i < len(m.Imports); i++ {
		providers = append(providers, m.Imports[i].Providers...)
	}

	if m.Name == "" {
		m.Name = "\"Unnamed\""
	}
	logger.Success(fmt.Sprintf("%s module successfully registered", m.Name))
	return routes, providers
}
