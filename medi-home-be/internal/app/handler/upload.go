package handler

import (
	"log"
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

func (h *UploadHandler) SingleUpload(c *gin.Context) {
	url, err := h.CloudinaryService.HandelRequest(c)
	log.Print("url", url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Upload failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *UploadHandler) MultiUpload(c *gin.Context) {
	log.Println(">> Nhận request multi_upload")

	urls, err := h.CloudinaryService.UploadMultiFileHandleRequestFromGin(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Upload failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upload successful",
		"urls":    urls,
	})
}
