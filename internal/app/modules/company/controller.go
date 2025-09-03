package company

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller  {
    c := deuterium.NewController("/company") 
    
    // Get company data
    c.Get("/").Register(func(ctx deuterium.Context) {})

    // Update comany data
    c.Put("/").Register(func(ctx deuterium.Context) {})

    return c
}
