package handler

import (
	"log"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	service service.CartService
}

func NewCartHandler(service service.CartService) *CartHandler {
	return &CartHandler{service}
}

func (h *CartHandler) GetCartUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	log.Print(id)
	medicine, err := h.service.GetCartByUser(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}
