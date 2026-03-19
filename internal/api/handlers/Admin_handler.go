package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/dto"
	"portfolio_backend_go/internal/service"

	"github.com/gin-gonic/gin"
)


type Admin_handler struct {

	Admin_Service *service.Admin_Service
}

func New_Admin_handler(Admin_service *service.Admin_Service) *Admin_handler {
	return &Admin_handler{Admin_Service: Admin_service}
}

func (h *Admin_handler) Admin_Login(c * gin.Context){
	var data dto.Admin_login_req

	if err := c.ShouldBindJSON(&data); err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Email and Password are required"})
		return
	}

	err := h.Admin_Service.Login(c.Request.Context(),data.Email,data.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
        "message": "Verify OTP",
    })


}
func (h *Admin_handler) Logout(c *gin.Context) {
    c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetSameSite(http.SameSiteNoneMode)

    c.JSON(http.StatusOK, gin.H{
        "message": "Logged out successfully",
    })
}

func (h *Admin_handler) CheckAuth(c *gin.Context) {
    uid, _ := c.Get("user_id")
    
    c.JSON(http.StatusOK, gin.H{
        "message": "auth working",
        "uid":     uid,
    })
}