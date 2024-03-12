package handler

import (
	"assignment-2/handler/order"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(orderHandler *order.Handler) {
	r := gin.Default()

	r.GET("/orders", orderHandler.GetAllOrders)
	r.POST("/orders", orderHandler.CreateOrders)
	r.PUT("/orders/:orderId", orderHandler.GetAllOrders)
	r.DELETE("/orders/:orderId", orderHandler.GetAllOrders)

	r.Run(":8000")
}
