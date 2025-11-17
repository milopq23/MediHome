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

// func (h *CartHandler) GetCartUser(c *gin.Context) {
// 	// claimsRaw, exists := c.Get("claims")
// 	// if !exists {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
// 	// 	return
// 	// }
// 	// claims := claimsRaw.(*util.Claims)
// 	// cartItem, err := h.service.GetCartByUser(int64(claims.UserID))
// 	// if err != nil {
// 	// 	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
// 	// 	return
// 	// }

//		id, _ := strconv.Atoi(c.Param("id"))
//		cartItem, err := h.service.GetCartByUser(int64(id))
//		if err != nil {
//			c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{
//			"data": cartItem,
//		})
//	}
func (h *CartHandler) GetCartUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cartItems, err := h.service.GetCartByUser(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	var total float64
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Price
	}

	response := dto.CartResponseDTO{
		CartItems:   cartItems,
		TotalAmount: total,
	}

	c.JSON(http.StatusOK, response)
}

// func (h *CartHandler) AddCart(c *gin.Context) {
// 	// claimsRaw, exists := c.Get("claims")
// 	// if !exists {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "No claims found"})
// 	// 	return
// 	// }
// 	// claims := claimsRaw.(*util.Claims)

// }

func (h *CartHandler) DeleteItemCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cartItem := h.service.DeleteCart(int64(id))
	c.JSON(http.StatusOK, gin.H{
		"DeleteCart": cartItem,
	})
}

func (h *CartHandler) AddCart(c *gin.Context) {
	var req dto.CartRequestDTO

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

	req = dto.CartRequestDTO{
		CartID:     cart.CartID,
		MedicineID: req.MedicineID,
		Quantity:   req.Quantity,
		SelectType: req.SelectType,
	}

	if _, err := h.service.AddMedicineToCart(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thêm thành công"})
}

func (h *CartHandler) UpdateCart(c *gin.Context) {
	var req dto.CartRequestDTO
	cartitem_id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req = dto.CartRequestDTO{
		// CartID:     cart.CartID,
		CartItemID: int64(cartitem_id),
		// MedicineID: req.MedicineID,
		Quantity:   req.Quantity,
		SelectType: req.SelectType,
	}
	if _, err := h.service.UpdateCart(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Cart"})
}
