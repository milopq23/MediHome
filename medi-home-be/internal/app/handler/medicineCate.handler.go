package handlers

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicineCateHandler struct {
	service service.MedicineCateService
}

func NewMedicineCateHandler(service service.MedicineCateService) *MedicineCateHandler {
	return &MedicineCateHandler{service}
}

func (h *MedicineCateHandler) GetAll(c *gin.Context) {
	medicineCates, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medicineCates)
}

func (h *MedicineCateHandler) ListChildren(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	medicineCate, err := h.service.ListChildren(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine category not found"})
		return
	}
	c.JSON(http.StatusOK, medicineCate)
}

func (h *MedicineCateHandler) Create(c *gin.Context) {
	var medicineCate model.MedicineCate
	if err := c.ShouldBindJSON(&medicineCate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMedicineCate, err := h.service.Create(medicineCate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMedicineCate)
}
func (h *MedicineCateHandler) Patch(c *gin.Context) {
	var medicineCate model.MedicineCate
	if err := c.ShouldBindJSON(&medicineCate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedMedicineCate, err := h.service.Patch(medicineCate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMedicineCate)
}

func (h *MedicineCateHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Medicine category deleted successfully"})
}
