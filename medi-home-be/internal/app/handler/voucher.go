package handler

import (
	"log"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VoucherHandler struct {
	service service.VoucherService
}

func NewVoucherHandler(service service.VoucherService) *VoucherHandler {
	return &VoucherHandler{service}
}

func (h *VoucherHandler) GetVoucherByCode(c *gin.Context) {
	code := c.Param("code")
	voucher, err := h.service.GetVoucherByCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func (h *VoucherHandler) FindActiveVoucher(c *gin.Context) {
	totalStr := c.Query("total")
	total, err := strconv.ParseFloat(totalStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total value"})
		return
	}
	voucher, err := h.service.FindVoucherActive(float64(total))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func (h *VoucherHandler) GetAllVoucher(c *gin.Context) {
	voucher, err := h.service.GetAllVoucher()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func (h *VoucherHandler) GetDetailVoucher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	voucher, err := h.service.GetDetailVoucher(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, voucher)
}

func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var voucher model.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdVoucher, err := h.service.CreateVoucher(voucher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, createdVoucher)
}

func (h *VoucherHandler) PatchVoucher(c *gin.Context) {
	voucher_id, _ := strconv.Atoi(c.Param("id"))

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(input)
	updatedVoucher, err := h.service.UpdateVoucher(int64(voucher_id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedVoucher)
}

func (h *VoucherHandler) DeleteVoucher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteVoucher(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
