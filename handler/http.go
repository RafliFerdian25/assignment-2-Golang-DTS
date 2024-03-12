package handler

import (
	"assignment-2/handler/order"

	"github.com/gin-gonic/gin"
)

func NewHttpServer(orderHandler *order.Handler) {
	r := gin.Default()

	r.GET("/orders", orderHandler.GetAllOrders)
	r.POST("/orders", orderHandler.CreateOrder)
	r.PUT("/orders/:orderId", orderHandler.UpdateOrder)
	r.DELETE("/orders/:orderId", orderHandler.DeleteOrder)

	r.Run(":8000")
}
