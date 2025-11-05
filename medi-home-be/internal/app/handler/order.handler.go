package handler

import (
	"medi-home-be/internal/app/service"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

// func (h *OrderHandler) CheckOut(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var orderRequest dto.OrderRequestDTO
// 	if err := c.ShouldBindJSON(&orderRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	order, err := h.service.GetCartByUser(id)

// }
