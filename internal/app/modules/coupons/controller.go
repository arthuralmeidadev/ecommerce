package coupons

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/coupons")

	// Get coupons
	c.Get("/").Register(func(ctx deuterium.Context) {})

	// Create coupons
	c.Post("/").Register(func(ctx deuterium.Context) {})

	// Get coupon
	c.Get("/:id").Register(func(ctx deuterium.Context) {})

	// Award coupon
	c.Post("/:id/award").Register(func(ctx deuterium.Context) {})

	// Redeem coupon
	c.Post("/:id/redeem").Register(func(ctx deuterium.Context) {})

	// Delete coupons
	c.Delete("/:id").Register(func(ctx deuterium.Context) {})

	return c
}
