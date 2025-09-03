package customers

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Customers",
		Controller: controller(),
	}
}
