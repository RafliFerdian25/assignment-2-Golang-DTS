package order

import (
	"assignment-2/service/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *order.Service
}

func NewHandler(svc *order.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) GetAllOrders(ctx *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, orders)
}
