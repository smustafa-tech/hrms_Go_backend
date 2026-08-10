package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/config"
	"github.com/smustafa-tech/hrms-backend/internal/handler"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
	"github.com/smustafa-tech/hrms-backend/internal/routes"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	router := gin.Default()

	// CORS — allow all frontend origins
	allowedOrigins := map[string]bool{
		"http://localhost:3000":                                 true, // Docker local
		"http://localhost:5173":                                 true, // Vite dev
		"https://hrms-frontend-production-c37d.up.railway.app": true, // Railway old
		"https://hrms-frontend-production-f099.up.railway.app": true, // Railway new
	}

	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
		}
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization,x-organization-slug,slug")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Wire auth
	authRepo := repository.NewAuthRepository(config.DB)
	authSvc := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	// Wire employee
	empRepo := repository.NewEmployeeRepository(config.DB)
	empSvc := service.NewEmployeeService(empRepo)
	empHandler := handler.NewEmployeeHandler(empSvc)

	// Wire attendance
	attRepo := repository.NewAttendanceRepository(config.DB)
	attSvc := service.NewAttendanceService(attRepo)
	attHandler := handler.NewAttendanceHandler(attSvc)

	api := router.Group("/api")
	routes.RegisterAuthRoutes(api, authHandler)
	routes.RegisterProtectedRoutes(api, empHandler, attHandler)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("✅ Server started at :" + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
