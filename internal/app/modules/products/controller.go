package products

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/products")

	// Get products
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Register product
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get product
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update product
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Deregister product
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
