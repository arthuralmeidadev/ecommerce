package orders

import (
	"ecommerce/internal/app/modules/test"
	"ecommerce/pkg/deuterium"
)

func controller() deuterium.Controller {
	c := deuterium.NewController("/orders")

	// Get Orders
	c.Get("/").Register(func(ctx deuterium.Context) {
		logger := &deuterium.Logger{Context: "Orders"}
		c := deuterium.GetContainer()
		var p test.TestProvider
		if err := c.Inject(&p); err != nil {
			logger.Error(err.Error())
			ctx.Response().InternalServerError("Internal Server Error")
			return
		}
		p.Test()
	})

	c.Post("/").Register(func(ctx deuterium.Context) {})

	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	c.Post("/:id/approve").Register(func(ctx deuterium.Context) {})

	c.Post("/:id/deny").Register(func(ctx deuterium.Context) {})

	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
