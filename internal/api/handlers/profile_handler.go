package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/models"
	"portfolio_backend_go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Profile_handler struct {
	Service *service.Profile_Service
}

func New_Profile_handler(s *service.Profile_Service) *Profile_handler {
	return &Profile_handler{Service: s}
}

func (h *Profile_handler) GetProfile(c *gin.Context) {
	profile, err := h.Service.GetProfile(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *Profile_handler) UpdateProfile(c *gin.Context) {
	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateProfile(c.Request.Context(), profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// ─── Skill Categories ─────────────────────────────────────────────────────────

func (h *Profile_handler) AddSkillCategory(c *gin.Context) {
	var cat models.SkillCategory
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddSkillCategory(c.Request.Context(), cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add skill category"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill category added"})
}

func (h *Profile_handler) UpdateSkillCategory(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("catIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category index"})
		return
	}
	var cat models.SkillCategory
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateSkillCategory(c.Request.Context(), idx, cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill category updated"})
}

func (h *Profile_handler) DeleteSkillCategory(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("catIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category index"})
		return
	}
	if err := h.Service.DeleteSkillCategory(c.Request.Context(), idx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill category deleted"})
}

// ─── Skills inside a Category ─────────────────────────────────────────────────

func (h *Profile_handler) AddSkillToCategory(c *gin.Context) {
	catIdx, err := strconv.Atoi(c.Param("catIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category index"})
		return
	}
	var skill models.SkillItem
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddSkillToCategory(c.Request.Context(), catIdx, skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill added to category"})
}

func (h *Profile_handler) UpdateSkillInCategory(c *gin.Context) {
	catIdx, err := strconv.Atoi(c.Param("catIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category index"})
		return
	}
	skillIdx, err := strconv.Atoi(c.Param("skillIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill index"})
		return
	}
	var skill models.SkillItem
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateSkillInCategory(c.Request.Context(), catIdx, skillIdx, skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill updated"})
}

func (h *Profile_handler) DeleteSkillFromCategory(c *gin.Context) {
	catIdx, err := strconv.Atoi(c.Param("catIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category index"})
		return
	}
	skillIdx, err := strconv.Atoi(c.Param("skillIndex"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill index"})
		return
	}
	if err := h.Service.DeleteSkillFromCategory(c.Request.Context(), catIdx, skillIdx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Skill deleted"})
}

// ─── Contacts / Socials ───────────────────────────────────────────────────────

func (h *Profile_handler) AddContact(c *gin.Context) {
	var input models.Social
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.AddContact(c.Request.Context(), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add contact"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact added successfully"})
}

func (h *Profile_handler) UpdateContact(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	var input models.Social
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.UpdateContact(c.Request.Context(), idx, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact updated successfully"})
}

func (h *Profile_handler) DeleteContact(c *gin.Context) {
	idx, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid index"})
		return
	}
	if err := h.Service.DeleteContact(c.Request.Context(), idx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
