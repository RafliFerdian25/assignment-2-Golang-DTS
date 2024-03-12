package main

import (
	"assignment-2/database"
	"assignment-2/handler"
	orderHandler "assignment-2/handler/order"
	"assignment-2/repo/order"
	orderService "assignment-2/service/order"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	// Setup Repo
	orderRepo := order.NewOrderRepository(db)

	// Setup Service
	orderService := orderService.NewService(orderRepo)

	// Setup Handler
	orderHandler := orderHandler.NewHandler(orderService)

	handler.NewHttpServer(orderHandler)
}
