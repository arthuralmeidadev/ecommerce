package customers

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/customers")

	// Get customers
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Create customer
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get customer
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update customer
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Delete customer
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	// Get customer's orders
	c.Get("/:id/orders").Register(func(ctx deuterium.Context) {})

	// Get customer's coupons
	c.Get("/:id/coupons").Register(func(ctx deuterium.Context) {})

	// Get customer's payments
	c.Get("/:id/payments").Register(func(ctx deuterium.Context) {})

	// Get customer's return appeals
	c.Get("/:id/returns").Register(func(ctx deuterium.Context) {})

	return c
}
