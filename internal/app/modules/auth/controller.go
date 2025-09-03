package auth

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller {
	c := deuterium.NewController("/auth")

	// Login
	c.Post("/").Register(func(ctx deuterium.Context) {
		var body AuthRequest
		ctx.Request().Body(&body)
	})

	// Logout
	c.Delete("/").Register(func(ctx deuterium.Context) {
	})

	return c
}
