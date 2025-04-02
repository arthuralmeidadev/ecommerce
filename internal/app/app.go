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
	app.Get("app-test").UseMiddleware(func(ctx deuterium.Context) {
		deuterium.GetLogger().Success("MIDDLEWARE 1")
        ctx.Next()
	}).Register(func(ctx deuterium.Context) {
		deuterium.GetLogger().Debug("TEST APP")
	})
	app.Listen("", port)
}
