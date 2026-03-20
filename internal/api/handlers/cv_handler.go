package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type CV_Handler struct {
	Service *service.CV_Service
}

func New_CV_Handler(s *service.CV_Service) *CV_Handler {
	return &CV_Handler{Service: s}
}

func (h *CV_Handler) UploadCV(c *gin.Context) {
	file, header, err := c.Request.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV file is required"})
		return
	}
	defer file.Close()

	// remove strict content-type check — browsers sometimes send
	// "application/octet-stream" for PDFs, so check by extension instead
	if !strings.HasSuffix(strings.ToLower(header.Filename), ".pdf") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only PDF files are allowed"})
		return
	}

	if header.Size > 5*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV must be under 5MB"})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		name = header.Filename
	}

	if err := h.Service.UploadCV(c.Request.Context(), file, header, name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload CV"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "CV uploaded successfully!"})
}
func (h *CV_Handler) GetAllCVs(c *gin.Context) {
	cvs, err := h.Service.GetAllCVs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch CVs"})
		return
	}

	list := []interface{}{}
	for key, value := range cvs {
		cv, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		cv["id"] = key
		list = append(list, cv)
	}

	c.JSON(http.StatusOK, list)
}

func (h *CV_Handler) SetActiveCV(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV ID is required"})
		return
	}

	if err := h.Service.SetActiveCV(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set active CV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Active CV updated!"})
}

func (h *CV_Handler) DeleteCV(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CV ID is required"})
		return
	}

	if err := h.Service.DeleteCV(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete CV"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CV deleted!"})
}

func (h *CV_Handler) GetActiveCV(c *gin.Context) {
	cv, err := h.Service.GetActiveCV(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch active CV"})
		return
	}
	if cv == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No active CV found"})
		return
	}

	c.JSON(http.StatusOK, cv)
}