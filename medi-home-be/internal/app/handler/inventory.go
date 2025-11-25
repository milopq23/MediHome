package handler

import (
	"log"
	"medi-home-be/internal/app/model"
	"medi-home-be/internal/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	service service.InventoryService
}

func NewInventoryHandler(service service.InventoryService) *InventoryHandler {
	return &InventoryHandler{service}
}

func (h *InventoryHandler) GetAll(c *gin.Context) {
	inventories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventories)
}

func (h *InventoryHandler) FindMedicine(c *gin.Context) {
	medicine_id, err := strconv.Atoi(c.Param("id"))

	inventories, err := h.service.FindMedicine(int64(medicine_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	batchselling, err := h.service.DetailBatchSelling(int64(medicine_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := struct {
		Inventories interface{} `json:"inventories"`
		InventoryID int64       `json:"inventory_id"`
	}{
		Inventories: inventories,
		InventoryID: batchselling.InventoryID,
	}
	c.JSON(http.StatusOK, response)
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	inventories, err := h.service.GetInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inventories)
}

func (h *InventoryHandler) Create(c *gin.Context) {
	var inventory model.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newInventory, err := h.service.Create(inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newInventory)
}

func (h *InventoryHandler) UpdateBatchSelling(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var inventory struct {
		InventoryID int64 `json:"inventory_id"`
	}
	log.Print(inventory.InventoryID)
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(inventory.InventoryID)

	updatedInventory, err := h.service.UpdateSellingBatch(int64(id), int64(inventory.InventoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedInventory)
}

func (h *InventoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
