package orders

import (
	"ecommerce/pkg/api"
)

func NewOrdersModule() *api.Module {
	return &api.Module{
		Name:              "Orders",
		ControllerFactory: &OrdersControllerFactory{},
		Providers:         []any{&OrdersProvider{}},
	}
}
