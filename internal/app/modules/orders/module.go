package orders

import (
	"ecommerce/internal/app/modules/test"
	"ecommerce/pkg/deuterium"
)

func Module() *deuterium.Module {
	return &deuterium.Module{
		Name: "Orders",
		Imports: []*deuterium.Module{
			test.Module(),
		},
		Controller: controller(),
		Providers:  []any{&OrdersProvider{}},
	}
}
