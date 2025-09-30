package handler

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShippingHandler struct {
	service service.ShippingService
}	

func NewShippingHandler(service service.ShippingService) *ShippingHandler {
	return &ShippingHandler{service}
}

func (h *ShippingHandler) GetAll(c *gin.Context) {
	shippings, err := h.service.GetAll()	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shippings)
}

// func (h *ShippingHandler) GetByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	shipping, err := h.service.GetByID(uint(id))	
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Shipping not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, shipping)
// }

func (h *ShippingHandler) Create(c *gin.Context) {
	var shipping model.Shipping
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newShipping, err := h.service.Create(shipping)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(http.StatusCreated, newShipping)
}

// func (h *ShippingHandler) Update(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var input model.Shipping
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	updated, err := h.service.Update(uint(id))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}	
// 	c.JSON(http.StatusOK, updated)
// }

// func (h *ShippingHandler) Patch(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	var input map[string]interface{}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	updated, err := h.service.Patch(int64(id), input)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, updated)
// }

func (h *ShippingHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))	
	if err := h.service.Delete(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
