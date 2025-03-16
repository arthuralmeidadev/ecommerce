package app

import (
	"ecommerce/pkg/api"
)

func Run() {
	port := 5000
	handler := MapRoutesToHandler() 
	server := api.NewServer("", port, handler)
	server.Listen()
}
