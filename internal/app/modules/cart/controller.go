package cart

import "ecommerce/pkg/deuterium"

func controller() deuterium.Controller  {
    c := deuterium.NewController("/cart") 

    // Create cart 
    c.Post("/").Register(func(ctx deuterium.Context) {})

    // Get cart 
    c.Get("/").Register(func(ctx deuterium.Context) {})

    // Add item to cart 
    c.Post("/item").Register(func(ctx deuterium.Context) {})

    // Update item from cart 
    c.Put("/item/:itemId").Register(func(ctx deuterium.Context) {})

    // Remove item from cart 
    c.Delete("/item/:itemId").Register(func(ctx deuterium.Context) {})

    // Clear cart 
    c.Delete("/").Register(func(ctx deuterium.Context) {})
    
    return c
}
