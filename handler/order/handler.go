package order

import (
	"assignment-2/core"
	"assignment-2/service/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	orderService *order.Service
}

func NewHandler(svc *order.Service) *Handler {
	return &Handler{
		orderService: svc,
	}
}

func (h *Handler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, orders)
}

func (h *Handler) CreateOrders(ctx *gin.Context) {
	// Binding request body to struct
	var OrderRequest core.OrderRequest
	if err := ctx.Bind(&OrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call orderService to create category
	order, err := h.orderService.CreateOrder(OrderRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail create category",
			"error":   err.Error(),
		})
		return
	}

	// Return response if success
	ctx.JSON(http.StatusOK, order)
}
