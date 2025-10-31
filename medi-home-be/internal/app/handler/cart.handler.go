package handler

import (
	"medi-home-be/internal/app/dto"
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
	// claimsRaw, exists := c.Get("claims")
	// if !exists {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
	// 	return
	// }
	// claims := claimsRaw.(*util.Claims)
	// cartItem, err := h.service.GetCartByUser(int64(claims.UserID))
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	// 	return
	// }

	id, _ := strconv.Atoi(c.Param("id"))
	cartItem, err := h.service.GetCartByUser(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": cartItem,
	})
}

// func (h *CartHandler) AddCart(c *gin.Context) {
// 	// claimsRaw, exists := c.Get("claims")
// 	// if !exists {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
// 	// 	return
// 	// }
// 	// claims := claimsRaw.(*util.Claims)

// }

func (h *CartHandler) AddCart(c *gin.Context) {
	var req dto.AddCartRequestDTO

	id, _ := strconv.Atoi(c.Param("id"))
	cart, err := h.service.GetCart(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.AddMedicineToCart(cart.CartID, req.MedicineID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Thêm thành công": item})
}
