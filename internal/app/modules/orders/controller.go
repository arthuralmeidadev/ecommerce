package orders

import (
	"ecommerce/pkg/deuterium"
)

func controller() deuterium.Controller {
	c := deuterium.NewController("/orders")

	// Get Orders
	c.Get("/").Register(func(ctx deuterium.Context) {
		deuterium.GetLogger().Debug("GET ORDERS")
	})

	// Place Order
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get Order
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update Order
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Approve Order
	c.Post("/:id/approve").Register(func(ctx deuterium.Context) {})

	// Deny Order
	c.Post("/:id/deny").Register(func(ctx deuterium.Context) {})

	// Cancel Order
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
