package cart

import "ecommerce/pkg/deuterium"

func Module() *deuterium.Module {
    return &deuterium.Module{
        Name: "Cart",
        Controller: controller(),
    }
}
