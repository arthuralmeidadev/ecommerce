package handlers

import (
	"ecommerce/pkg/api"
	"ecommerce/internal/pkg/dto"
)

func GetOrders(ctx api.Context) {

}

func PlaceOrder(ctx api.Context) {
	var order dto.OrderRequestDTO
	err := ctx.Request().Body(&order)
	if err != nil {
		ctx.Response().BadRequestError("Invalid request body!")
	}

	customerId := order.Customer.Id
	println(customerId)
}

func GetOrder(ctx api.Context) {
	ctx.Response().WriteStatusMessage(200, "success")
	return
}

func UpdateOrder(ctx api.Context) {

}

func ApproveOrder(ctx api.Context) {

}

func DenyOrder(ctx api.Context) {

}

func CancelOrder(ctx api.Context) {

}
