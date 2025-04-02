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
	if m.Controller != nil {
		routes = m.Controller.register()
		logger.Success(fmt.Sprintf("%s controller successfully registered", m.Name))
	}

	providers := m.Providers
	for i := 0; i < len(m.Imports); i++ {
		providers = append(providers, m.Imports[i].Providers...)
	}

	logger.Success(fmt.Sprintf("%s module successfully registered", m.Name))
	return routes, providers
}
