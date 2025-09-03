package payments

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Payments",
		Controller: controller(),
	}
}
