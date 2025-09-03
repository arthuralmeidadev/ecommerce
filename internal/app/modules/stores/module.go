package stores

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Stores",
		Controller: controller(),
	}
}
