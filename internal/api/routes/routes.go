package routes

import (
	"portfolio_backend_go/internal/api/handlers"
	"portfolio_backend_go/internal/api/middleware"

	"firebase.google.com/go/v4/auth" 
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, 
	projectHandler *handlers.ProjectHandler, 
	otpHandler *handlers.OTP_handler,
	adminHanler *handlers.Admin_handler,
	trackVisite *handlers.StatsHandler,
	authClient *auth.Client) {
    
    v1 := r.Group("/api/v1")
    {
        v1.GET("/projects", projectHandler.GetAll)
		v1.POST("/otp/verify",otpHandler.Verify_otp)
		v1.POST("/login",adminHanler.Admin_Login)
		v1.POST("/track/visit",trackVisite.Track_Visitor)
		v1.POST("/track/downloads",trackVisite.TrackDownload)
		v1.GET("/track/stats",trackVisite.GetStats)
		


        admin := v1.Group("/admin")
		admin.Use(middleware.Auth_Middleware(authClient))
        {
            
            admin.POST("/projects", projectHandler.CreateProject)
			admin.GET("/checkAuth",adminHanler.CheckAuth)
			admin.POST("/logout",adminHanler.Logout)

			
        }

		
    }
}
