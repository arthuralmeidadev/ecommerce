package campaigns

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name:       "Campaigns",
		Controller: controller(),
	}
}
