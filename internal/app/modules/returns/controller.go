package returns

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/returns")

	// Get return appeals
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Submit return appeals
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get return appeal
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update return appeal
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Approve return appeal
	c.Post("/:id/approve").Register(func(ctx deuterium.Context) {})

	// Deny return appeal
	c.Post("/:id/deny").Register(func(ctx deuterium.Context) {})

	// Cancel return appeal
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	// Confirm returned products
	c.Post("/:id/confirm").Register(func(ctx deuterium.Context) {})

	return c
}
