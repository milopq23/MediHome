package handler

import (
	"log"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	service service.AddressService
}

func NewAddressHandler(service service.AddressService) *AddressHandler {
	return &AddressHandler{service}
}

func (h *AddressHandler) AddAddress(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetAddressUser(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user_id not found"})
		return
	}
	log.Print(user.UserID)
	var address model.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := h.service.AddAddress(int64(user.UserID), address); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Thêm thành công"})
}
