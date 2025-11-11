package handler

import (
	"log"
	"medi-home-be/internal/app/dto"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{service}
}

func (h *OrderHandler) CheckOut(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var orderRequest dto.OrderRequestDTO
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(orderRequest)

	orderRequest = dto.OrderRequestDTO{
		UserID:        int64(id),
		VoucherCode:   orderRequest.VoucherCode,
		ShippingID:    orderRequest.ShippingID,
		PaymentMethod: orderRequest.PaymentMethod,
	}
	order, err := h.service.CheckOut(orderRequest)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed checkout"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Thanh toán thành công", "data": order})
}

func (h *OrderHandler) GetAllOrder(c *gin.Context) {
	orders, err := h.service.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy tất cả đơn hàng",
		"data":    orders,
	})

}
