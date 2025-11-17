package handler

import (
	"log"
	"medi-home-be/internal/app/dto"
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
	log.Print("medicine", medicine)
	newMedicine, err := h.service.Create(medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMedicine)
}

func (h *MedicineHandler) SaveImages(c *gin.Context) {
	medicineIDStr := c.Param("medicine_id")
	medicineID, err := strconv.ParseInt(medicineIDStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid medicine_id"})
		return
	}
	log.Print("medicine_id",medicineID)

	var body dto.ImageMedicineDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid body"})
		return
	}
	log.Print("image",body.Urls)

	if err := h.service.SaveImages(medicineID, body.Urls); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Images saved successfully",
	})
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
	c.JSON(http.StatusOK, gin.H{"message": "Xoá thành công medicine"})
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

// func (h *MedicineHandler) DetailMedicine(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	medicine, err := h.service.DetailMedicine(int64(id))
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, medicine)
// }

func (h *MedicineHandler) DetailMedicineUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	medicine, err := h.service.DetailMedicineUser(int64(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}
