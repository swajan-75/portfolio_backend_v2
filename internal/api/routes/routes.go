package routes

import (
	"portfolio_backend_go/internal/api/handlers"
	"portfolio_backend_go/internal/api/middleware"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// add storageHandler param
func SetupRoutes(
	r *gin.Engine,
	projectHandler *handlers.ProjectHandler,
	otpHandler *handlers.OTP_handler,
	adminHanler *handlers.Admin_handler,
	trackVisite *handlers.StatsHandler,
	storageHandler *handlers.Storage_Handler, 
	authClient *auth.Client,
) {
	login_limiter := middleware.NewIPRateLimiter(rate.Limit(5.0/60.0), 3)
	global_limiter := middleware.NewIPRateLimiter(rate.Limit(100.0/60), 3)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.RateLimitMiddleware(global_limiter))
	{
		v1.GET("/projects", projectHandler.GetAll)
		v1.GET("/projects/:slug", projectHandler.GetProject)
		v1.POST("/otp/verify", middleware.RateLimitMiddleware(login_limiter), otpHandler.Verify_otp)
		v1.POST("/login", middleware.RateLimitMiddleware(login_limiter), adminHanler.Admin_Login)

		admin := v1.Group("/admin")
		admin.Use(middleware.Auth_Middleware(authClient))
		{
			admin.POST("/projects", projectHandler.CreateProject)
			admin.PUT("/projects/:slug", projectHandler.UpdateProject)
			admin.DELETE("/projects/:slug", projectHandler.DeleteProject)
			admin.POST("/upload/image", storageHandler.UploadImage) // ← add this

			admin.GET("/checkAuth", adminHanler.CheckAuth)
			admin.POST("/logout", adminHanler.Logout)
			admin.POST("/track/visit", trackVisite.Track_Visitor)
			admin.POST("/track/downloads", trackVisite.TrackDownload)
			admin.GET("/track/stats", trackVisite.GetStats)
		}
	}
}