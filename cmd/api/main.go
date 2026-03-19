package main

import (
	"context"
	"log"
	"portfolio_backend_go/internal/api/handlers"
	"portfolio_backend_go/internal/api/routes"
	"portfolio_backend_go/internal/config"
	"portfolio_backend_go/internal/repository"
	"portfolio_backend_go/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, fetching variables from system environment")
	}

	// 1. Init Firebase
	app := config.Init_firebase()

	// 2. Init Database
	dbClient, err := app.Database(ctx)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// 3. Init Auth
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error initializing auth: %v", err)
	}

	// 4. Dependency Injection
	projectRepo    := repository.New_Project_repo(dbClient)
	projectService := service.New_Project_Service(projectRepo)
	projectHandler := handlers.NewProjectHandler(projectService)

	otpRepo    := repository.New_OTP_repo(dbClient)
	otpService := service.New_Otp_service(otpRepo)
	otpHandler := handlers.New_Otp_handler(otpService)

	adminService := service.New_Admin__Service(otpService)
	adminHandler := handlers.New_Admin_handler(adminService)

	trackVisite := handlers.NewStatsHandler(dbClient)

	storageRepo    := repository.New_Storage_repo()
	storageService := service.New_Storage_Service(storageRepo)
	storageHandler := handlers.New_Storage_Handler(storageService)

    cvRepo      := repository.New_CV_repo(dbClient)
    cvService   := service.New_CV_Service(cvRepo, storageRepo)
    cvHandler   := handlers.New_CV_Handler(cvService)

	// 5. Gin Setup
	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://swajan.vercel.app")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 6. Routes
	routes.SetupRoutes(server, projectHandler, otpHandler, adminHandler, trackVisite, storageHandler,cvHandler, authClient)

	server.Run(":8000")
}