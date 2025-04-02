package app

import (
	"ecommerce/internal/app/modules/orders"
	"ecommerce/pkg/deuterium"
)

func Run() {
	port := 5000
	modules := []*deuterium.Module{
		orders.Module(),
	}
	app := deuterium.NewApp(modules)
	var middlewares []deuterium.ContextHandler
	middlewares = append(
		middlewares,
		func(ctx deuterium.Context) {
			deuterium.GetLogger().Debug("MIDDLEWARE 1")
			ctx.Next()
		},
        func(ctx deuterium.Context) {
			deuterium.GetLogger().Debug("MIDDLEWARE 2")
			ctx.Next()
        },
	)
	app.Get("app-test").UseMiddlewares(middlewares).Register(
		func(ctx deuterium.Context) {
			deuterium.GetLogger().Debug("TEST APP")
		},
	)
	app.Listen("", port)
}
