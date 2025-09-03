package stores

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/stores")

	// Get stores
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Register store
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get store
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update store
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Deregister store
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	// Get store orders
	c.Get("/:id/orders").Register(func(ctx deuterium.Context) {})

	// Get store coupons
	c.Get("/:id/coupons").Register(func(ctx deuterium.Context) {})

	// Get store payments
	c.Get("/:id/payments").Register(func(ctx deuterium.Context) {})

	// Get store products
	c.Get("/:id/products").Register(func(ctx deuterium.Context) {})

	// Get store return appeals
	c.Get("/:id/returns").Register(func(ctx deuterium.Context) {})

	return c
}
