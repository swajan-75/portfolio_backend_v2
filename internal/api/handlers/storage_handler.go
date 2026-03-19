package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type Storage_Handler struct {
	Service *service.Storage_Service
}

func New_Storage_Handler(s *service.Storage_Service) *Storage_Handler {
	return &Storage_Handler{Service: s}
}

func (h *Storage_Handler) UploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}
	defer file.Close()

	// validate file type
	contentType := header.Header.Get("Content-Type")
	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/webp" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPEG, PNG and WebP images are allowed"})
		return
	}

	// validate file size (2MB max)
	if header.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image must be under 2MB"})
		return
	}

	url, err := h.Service.UploadImage(c.Request.Context(), file, header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}