package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/dto"
	"portfolio_backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	Service *service.Project_Service
}

func NewProjectHandler(s *service.Project_Service) *ProjectHandler {
	return &ProjectHandler{Service: s}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req dto.Create_Project_dto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreateProject(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "project with this title already exists" {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "fail",
				"message": "Duplicate Entry: A project with this title already exists.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project added successfully!"})
}

func (h *ProjectHandler) GetAll(c *gin.Context) {
	projects, err := h.Service.GetProjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	projectList := []interface{}{}
	for _, value := range projects {
		projectList = append(projectList, value)
	}

	c.JSON(http.StatusOK, projectList)
}

func (h *ProjectHandler) GetProject(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project slug is required"})
		return
	}

	project, err := h.Service.GetProjectBySlug(c.Request.Context(), slug)
	if err != nil {
		if err.Error() == "project not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project slug is required"})
		return
	}

	var req dto.Update_Project_dto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.UpdateProject(c.Request.Context(), slug, req)
	if err != nil {
		if err.Error() == "project not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully!"})
}

func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Project slug is required"})
		return
	}

	err := h.Service.DeleteProject(c.Request.Context(), slug)
	if err != nil {
		if err.Error() == "project not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully!"})
}