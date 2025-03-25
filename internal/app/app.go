package app

import (
	"ecommerce/internal/app/modules/orders"
	"ecommerce/pkg/api"
)

func Run() {
	port := 5000
	modules := make([]*api.Module, 0)
	modules = append(modules, orders.NewOrdersModule())
	app := api.NewApp(modules)
	app.Listen("", port)
}
