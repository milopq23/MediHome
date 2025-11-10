package handler

import (
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DosageFormHandler struct{
	service service.DosageFormService
}

func NewDosageFormHandler(service service.DosageFormService) *DosageFormHandler{
	return &DosageFormHandler{service}
}

func (h *DosageFormHandler) GetAll(c *gin.Context){
	dosages, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dosages)
}

func (h *DosageFormHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	dosage,err := h.service.GetByID(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DosageForm not found"})
		return
	}
	c.JSON(http.StatusOK, dosage)
}

func (h *DosageFormHandler) Create(c *gin.Context) {
	var dosage model.DosageForm
	if err := c.ShouldBindJSON(&dosage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newDosage, err := h.service.Create(dosage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newDosage)
}

func (h *DosageFormHandler) Patch(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := h.service.Patch(int64(id), input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *DosageFormHandler) Delete(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.Delete(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted"})
}