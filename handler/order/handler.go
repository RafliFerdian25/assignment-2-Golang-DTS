package order

import (
	"assignment-2/core"
	"assignment-2/service/order"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (h *Handler) CreateOrder(ctx *gin.Context) {
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

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	// Get orderId from path and convert to uint
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	err = h.orderService.DeleteOrder(uint(orderId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail delete order",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success delete",
	})
}

func (h *Handler) UpdateOrder(ctx *gin.Context) {
	// Get orderId from path and convert to uint
	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order id",
		})
		return
	}

	// Binding request body to struct
	var OrderRequest core.OrderRequest
	if err := ctx.Bind(&OrderRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call orderService to update order
	order, err := h.orderService.UpdateOrder(uint(orderId), OrderRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "order not found",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "fail update order",
			"error":   err.Error(),
		})
		return
	}

	// Return response if success
	ctx.JSON(http.StatusOK, order)
}
