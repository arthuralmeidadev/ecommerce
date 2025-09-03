package payments

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/payments")

	// Get payments
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Create payment
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get payment
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Update payment
	c.Put("/:id").Register(func(ctx deuterium.Context) {})

	// Cancel payment
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	// Authorize credit card payment
	c.Post("/credit-card/authorize").Register(func(ctx deuterium.Context) {})

	// Capture credit card payment
	c.Post("/credit-card/capture").Register(func(ctx deuterium.Context) {})

	// Execute debit card payment
	c.Post("/debit-card").Register(func(ctx deuterium.Context) {})

	// Digital wallet payment
	c.Post("/digital-wallet").Register(func(ctx deuterium.Context) {})

	// Bank transfer payment
	c.Post("/bank-transfer").Register(func(ctx deuterium.Context) {})

	return c
}
