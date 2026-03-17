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

    // 1. Bind and Validate JSON
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 2. Call Service
    err := h.Service.CreateProject(c.Request.Context(), req)
    if err != nil {
        // 🚀 Check if the error is our "Duplicate" error
        if err.Error() == "project with this title already exists" {
            c.JSON(http.StatusConflict, gin.H{
                "status":  "fail",
                "message": "Duplicate Entry: A project with this title already exists.",
            })
            return
        }

        // For any other error (like database connection issues)
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