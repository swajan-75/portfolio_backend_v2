package handlers

import (
	"net/http"
	"portfolio_backend_go/internal/dto"
	"portfolio_backend_go/internal/service"

	"github.com/gin-gonic/gin"
)

type OTP_handler struct {
	Otp_service *service.Otp_service
}

func New_Otp_handler(otp_service *service.Otp_service) *OTP_handler {
	return &OTP_handler{Otp_service: otp_service}
}


func (s *OTP_handler) Verify_otp (c *gin.Context){
	var data dto.Otp_verify

	if err := c.ShouldBindJSON(&data); err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Email and OTP are required"})
		return
	}

	idToken,err := s.Otp_service.Verify_otp(c.Request.Context(),data.Email,data.OTP)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
            "error": err.Error(),
        })
        return
	}
	if idToken ==""{
		c.JSON(http.StatusUnauthorized, gin.H{
            "error": "IdToken invalid",
        })
        return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access_token",idToken,3600,"/","",true,true)

	c.JSON(http.StatusOK, gin.H{
        "message": "OTP Verified successfully. Welcome Admin.",
        "status":  "verified",
		"token" : idToken,
    })

}