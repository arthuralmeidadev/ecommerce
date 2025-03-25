package orders

import (
	"ecommerce/internal/app/modules/test"
	"ecommerce/pkg/api"
	"log"
)

func getOrders(ctx api.Context) {
	logger := &api.Logger{Context: "Orders"}
	c := api.GetContainer()
	var p test.TestProvider
	if err := c.Inject(&p); err != nil {
		log.Printf("Error %v", err)
		ctx.Response().InternalServerError("Internal Server Error")
		return
	}

	logger.Info("Test")
	logger.Warn("Test")
	logger.Debug("Test")
	logger.Success("Test")
	logger.Error("Test")
	logger.Fatal("Test")
}

func placeOrder(ctx api.Context) {}

func getOrder(ctx api.Context) {}

func updateOrder(ctx api.Context) {}

func approveOrder(ctx api.Context) {}

func denyOrder(ctx api.Context) {}

func cancelOrder(ctx api.Context) {}

type OrdersControllerFactory struct{}

func (f *OrdersControllerFactory) Make() api.Controller {
	c := api.NewController("/orders")
	c.Get("/").Register(getOrders)
	c.Post("/").Register(placeOrder)
	c.Get("/:id").Register(getOrder)
	c.Put("/:id").Register(updateOrder)
	c.Post("/:id/approve").Register(approveOrder)
	c.Post("/:id/deny").Register(denyOrder)
	c.Delete("/:id").Register(cancelOrder)
	return c
}
