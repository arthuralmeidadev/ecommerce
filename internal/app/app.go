package app

import (
	"ecommerce/internal/app/handlers"
	"ecommerce/internal/pkg/api"
)

func Run() {
	r := api.NewRouter()

	r.Get("/auth").Register(handlers.Authenticate)
	r.Delete("/auth").Register(handlers.TerminateSession)

	r.Get("/campaigns").Register(handlers.GetCampaigns)
	r.Post("/campaigns").Register(handlers.CreateCampaign)
	r.Get("/campaigns/:id").Register(handlers.GetCampaign)
	r.Put("/campaigns/:id").Register(handlers.UpdateCampaign)
	r.Delete("/campaigns/:id").Register(handlers.DeleteCampaign)

	r.Post("/cart").Register(handlers.CreateCart)
	r.Get("/cart").Register(handlers.GetCart)
	r.Post("/cart/item").Register(handlers.AddItemToCart)
	r.Put("/cart/item").Register(handlers.UpdateItemFromCart)
	r.Delete("/cart/item").Register(handlers.RemoveItemFromCart)
	r.Delete("/cart").Register(handlers.ClearCart)

	r.Get("/company").Register(handlers.GetCompanyInfo)
	r.Put("/company").Register(handlers.UpdateCompanyInfo)

	r.Get("/coupons").Register(handlers.GetCoupons)
	r.Post("/coupons").Register(handlers.RegisterCoupon)
	r.Get("/coupons/:id").Register(handlers.GetCoupon)
	r.Post("/coupons/:id/award").Register(handlers.AwardCoupon)
	r.Post("/coupons/:id/redeem").Register(handlers.RedeemCoupon)
	r.Delete("/coupons/:id").Register(handlers.DeleteCoupon)

	r.Get("/customers").Register(handlers.GetCustomers)
	r.Post("/customers").Register(handlers.CreateCustomer)
	r.Get("/customers/:id").Register(handlers.GetCustomer)
	r.Put("/customers/:id").Register(handlers.UpdateCustomer)
	r.Delete("/customers/:id").Register(handlers.DeleteCustomer)
	r.Get("/customers/:id/orders").Register(handlers.GetCustomers)
	r.Get("/customers/:id/coupons").Register(handlers.GetCoupons)
	r.Get("/customers/:id/payments").Register(handlers.GetCustomerPayments)
	r.Get("/customers/:id/returns").Register(handlers.GetCustomerReturnAppeals)
	r.Get("/customers/:id/transactions").Register(handlers.GetCustomerTransactions)

	r.Get("/orders").Register(handlers.GetOrders)
	r.Post("/orders").Register(handlers.PlaceOrder)
	r.Get("/orders/:id").Register(handlers.GetOrder)
	r.Put("/orders/:id").Register(handlers.UpdateOrder)
	r.Post("/orders/:id/approve").Register(handlers.ApproveOrder)
	r.Post("/orders/:id/deny").Register(handlers.DenyOrder)
	r.Delete("/orders/:id").Register(handlers.CancelOrder)

	r.Get("/payments").Register(handlers.GetPayments)
	r.Post("/payments").Register(handlers.CreatePayment)
	r.Get("/payments/:id").Register(handlers.GetPayment)
	r.Put("/payments/:id").Register(handlers.UpdatePayment)
	r.Delete("/payments/:id").Register(handlers.CancelPayment)

	r.Get("/products").Register(handlers.GetProducts)
	r.Post("/products").Register(handlers.RegisterProduct)
	r.Get("/products/:id").Register(handlers.GetProduct)
	r.Put("/products/:id").Register(handlers.UpdateProduct)
	r.Delete("/products/:id").Register(handlers.RemoveProduct)

	r.Get("/returns").Register(handlers.GetReturnAppeals)
	r.Post("/returns").Register(handlers.SubmitReturnAppeal)
	r.Get("/returns/:id").Register(handlers.GetReturnAppeal)
	r.Put("/returns/:id").Register(handlers.UpdateReturnAppeal)
	r.Post("/returns/:id/approve").Register(handlers.ApproveReturnAppeal)
	r.Post("/returns/:id/deny").Register(handlers.DenyOrder)
	r.Delete("/returns/:id").Register(handlers.CancelReturnAppeal)
	r.Post("/returns/:id/confirm").Register(handlers.ConfirmReturnedProducts)

	r.Get("/stores").Register(handlers.GetStores)
	r.Post("/stores").Register(handlers.RegisterStore)
	r.Get("/stores/:id").Register(handlers.GetStore)
	r.Put("/stores/:id").Register(handlers.UpdateStore)
	r.Delete("/stores/:id").Register(handlers.RemoveStore)
	r.Get("/stores/:id/campaigns").Register(handlers.GetStoreCampaings)
	r.Get("/stores/:id/coupons").Register(handlers.GetStoreCoupons)
	r.Get("/stores/:id/orders").Register(handlers.GetStoreOrders)
	r.Get("/stores/:id/products").Register(handlers.GetStoreProducts)
	r.Get("/stores/:id/returns").Register(handlers.GetStoreReturnAppeals)

	r.Get("/transactions").Register(handlers.GetTransactions)
	r.Get("/transactions/:id").Register(handlers.GetTransaction)
	r.Delete("/transactions/:id").Register(handlers.CancelTransaction)
	r.Post("/transactions/credit-card").Register(handlers.ExecuteCreditCardTransaction)
	r.Post("/transactions/debit-card").Register(handlers.ExecuteDebitCardTransaction)
	r.Post("/transactions/digital-wallet").Register(handlers.ExecuteDigitalWalletTransaction)
	r.Post("/transactions/bank-transfer").Register(handlers.ExecuteDirectBankTransaction)

	r.Get("/users").Register(handlers.GetUsers)
	r.Post("/users").Register(handlers.RegisterUser)
	r.Get("/users/:id").Register(handlers.GetUser)
	r.Put("/users/:id").Register(handlers.UpdateUser)
	r.Delete("/users/:id").Register(handlers.DeleteUser)

	port := 5000
	handler := r.Register()
	server := api.NewServer("", port, handler)
	server.Listen()
}
