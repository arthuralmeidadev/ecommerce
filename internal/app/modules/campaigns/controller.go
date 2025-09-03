package campaigns

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/campaigns")

	// Get campaigns
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Create campaign
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get campaign
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update campaign
	c.Patch("/:id").Register(func(ctx deuterium.Context) {})

	// End campaign
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
