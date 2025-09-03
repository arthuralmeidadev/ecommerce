package auth

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Auth",
		Controller: controller(),
	}
}
