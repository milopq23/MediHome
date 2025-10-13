package handler

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicineHandler struct {
	service service.MedicineService
}

func NewMedicineHandler(service service.MedicineService) *MedicineHandler {
	return &MedicineHandler{service}
}

func (h *MedicineHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	pagination, err := h.service.GetAll(page, pageSize)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(http.StatusOK, pagination)
}

func (h *MedicineHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	medicine, err := h.service.GetByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}

func (h *MedicineHandler) Create(c *gin.Context) {
	var medicine model.Medicine	
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMedicine, err := h.service.Create(medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}			
	c.JSON(http.StatusCreated, newMedicine)
}

func (h *MedicineHandler) Patch(c *gin.Context) {
	var medicine model.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedMedicine, err := h.service.Patch(medicine)	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedMedicine)
}

func (h *MedicineHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MedicineHandler) ListMedicine(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	pagination, err := h.service.ListMedicine(page, pageSize)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(http.StatusOK, pagination)
}
