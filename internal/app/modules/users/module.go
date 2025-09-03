package users

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Users",
		Controller: controller(),
	}
}
