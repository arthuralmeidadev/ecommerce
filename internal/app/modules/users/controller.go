package users

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/users")

	// Get users
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Create user
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get user
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update user
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Delete user
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
