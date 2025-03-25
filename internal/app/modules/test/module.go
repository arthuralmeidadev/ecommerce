package test

import "ecommerce/pkg/api"

func NewTestModule() *api.Module {
	return &api.Module{
		Name:      "Test",
        Providers: []any{&TestProvider{ v: "test"}},
	}
}
