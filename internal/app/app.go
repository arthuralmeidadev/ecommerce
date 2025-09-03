package app

import (
	"ecommerce/internal/app/modules/auth"
	"ecommerce/internal/app/modules/campaigns"
	"ecommerce/internal/app/modules/cart"
	"ecommerce/internal/app/modules/company"
	"ecommerce/internal/app/modules/coupons"
	"ecommerce/internal/app/modules/customers"
	"ecommerce/internal/app/modules/orders"
	"ecommerce/internal/app/modules/payments"
	"ecommerce/internal/app/modules/products"
	"ecommerce/internal/app/modules/returns"
	"ecommerce/internal/app/modules/stores"
	"ecommerce/internal/app/modules/users"
	"ecommerce/pkg/deuterium"
)

type modules []*deuterium.Module

func Run() {
	app := deuterium.NewApp(
		modules{
			auth.Module(),
			campaigns.Module(),
			cart.Module(),
			company.Module(),
			coupons.Module(),
			customers.Module(),
			orders.Module(),
			payments.Module(),
			products.Module(),
			returns.Module(),
			stores.Module(),
			users.Module(),
		},
	)

	app.SetName("Cloud Commerce")
	app.SetVersion(1, 0, 0)

	port := 5000
	app.Listen("", port)
}
