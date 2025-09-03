package products

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Products",
		Controller: controller(),
	}
}
