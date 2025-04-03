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
	app.Use(func(ctx deuterium.Context) {
		deuterium.GetLogger().Debug("App level mid")
		ctx.Next()
	})
	app.Get("app-test").
		Use(func(ctx deuterium.Context) {
			deuterium.GetLogger().Debug("App route mid")
		}).
		Register(
			func(ctx deuterium.Context) {
				deuterium.GetLogger().Debug("TEST APP")
			},
		)
	app.Listen("", port)
}
