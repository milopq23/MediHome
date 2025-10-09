package handler

import (
	"medi-home-be/internal/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	CloudinaryService service.CloudinaryService
}

func NewUploadHandler(service service.CloudinaryService) *UploadHandler {
	return &UploadHandler{CloudinaryService: service}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	url, err := h.CloudinaryService.HandelResquest(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Upload failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"url":     url,
	})
}
