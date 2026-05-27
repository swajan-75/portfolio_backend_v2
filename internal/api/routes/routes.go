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
	cvHandler *handlers.CV_Handler,
	profileHandler *handlers.Profile_handler,
	authClient *auth.Client,
) {
	login_limiter := middleware.NewIPRateLimiter(rate.Limit(5.0/60.0), 5)
	global_limiter := middleware.NewIPRateLimiter(rate.Limit(200.0/60), 20)

	v1 := r.Group("/api/v1")
	v1.Use(middleware.RateLimitMiddleware(global_limiter))
	{
		v1.GET("/projects", projectHandler.GetAll)
		v1.GET("/cv/active", cvHandler.GetActiveCV)
		v1.GET("/profile", profileHandler.GetProfile)
		v1.GET("/projects/:slug", projectHandler.GetProject)
		v1.POST("/otp/verify", middleware.RateLimitMiddleware(login_limiter), otpHandler.Verify_otp)
		v1.POST("/login", middleware.RateLimitMiddleware(login_limiter), adminHanler.Admin_Login)
		v1.POST("/track/visit", trackVisite.Track_Visitor)
		v1.POST("/track/downloads", trackVisite.TrackDownload)

		admin := v1.Group("/admin")
		admin.Use(middleware.Auth_Middleware(authClient))
		{
			admin.POST("/projects", projectHandler.CreateProject)
			admin.PUT("/projects/:slug", projectHandler.UpdateProject)
			admin.DELETE("/projects/:slug", projectHandler.DeleteProject)
			admin.POST("/upload/image", storageHandler.UploadImage) 
			

			admin.GET("/checkAuth", adminHanler.CheckAuth)
			admin.POST("/logout", adminHanler.Logout)
			admin.GET("/track/stats", trackVisite.GetStats)

			admin.POST("/cv", cvHandler.UploadCV)
			admin.GET("/cv", cvHandler.GetAllCVs)
			admin.PUT("/cv/:id/active", cvHandler.SetActiveCV)
			admin.DELETE("/cv/:id", cvHandler.DeleteCV)

			admin.POST("/profile", profileHandler.UpdateProfile)

			// Skill Categories
			admin.POST("/profile/skill-categories", profileHandler.AddSkillCategory)
			admin.PUT("/profile/skill-categories/:catIndex", profileHandler.UpdateSkillCategory)
			admin.DELETE("/profile/skill-categories/:catIndex", profileHandler.DeleteSkillCategory)

			// Skills inside a Category
			admin.POST("/profile/skill-categories/:catIndex/skills", profileHandler.AddSkillToCategory)
			admin.PUT("/profile/skill-categories/:catIndex/skills/:skillIndex", profileHandler.UpdateSkillInCategory)
			admin.DELETE("/profile/skill-categories/:catIndex/skills/:skillIndex", profileHandler.DeleteSkillFromCategory)

			// Profile Contacts/Socials endpoints
			admin.POST("/profile/contacts", profileHandler.AddContact)
			admin.PUT("/profile/contacts/:index", profileHandler.UpdateContact)
			admin.DELETE("/profile/contacts/:index", profileHandler.DeleteContact)
		}
	}
}
