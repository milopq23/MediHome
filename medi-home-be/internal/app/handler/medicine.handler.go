package handlers

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
	medicines, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(http.StatusOK, medicines)
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
