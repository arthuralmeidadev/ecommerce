package test

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:      "Test",
		Providers: []any{&TestProvider{v: "test"}},
	}
}
