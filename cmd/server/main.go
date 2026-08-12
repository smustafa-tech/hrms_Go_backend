package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/config"
	"github.com/smustafa-tech/hrms-backend/internal/handler"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
	"github.com/smustafa-tech/hrms-backend/internal/routes"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

func parseAllowedOrigins(envValue string) map[string]bool {
	defaults := map[string]bool{
		"http://localhost:3000":                                true,
		"http://localhost:5173":                                true,
		"https://hrms-frontend-production-c37d.up.railway.app": true,
		"https://hrms-frontend-production-f099.up.railway.app": true,
		"https://hrms-frontend-production-039b.up.railway.app": true,
	}

	result := make(map[string]bool)
	for k, v := range defaults {
		result[k] = v
	}

	if envValue != "" {
		for _, origin := range strings.Split(envValue, ",") {
			trimmed := strings.TrimSpace(origin)
			if trimmed != "" {
				result[trimmed] = true
			}
		}
	}
	return result
}

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	router := gin.Default()

	// CORS — allow origins from env var (comma-separated), with defaults
	allowedOrigins := parseAllowedOrigins(config.GetEnv("ALLOWED_ORIGINS"))

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

	// Wire user and settings
	userRepo := repository.NewUserRepository(config.DB)
	userSvc := service.NewUserService(userRepo)
	companyRepo := repository.NewCompanyRepository(config.DB)
	companySvc := service.NewCompanyService(companyRepo)
	settingsRepo := repository.NewSettingsRepository(config.DB)
	settingsSvc := service.NewSettingsService(settingsRepo)
	userHandler := handler.NewUserHandler(userSvc, companySvc, settingsSvc)

	// Wire notifications
	notifRepo := repository.NewNotificationRepository(config.DB)
	notifSvc := service.NewNotificationService(notifRepo)
	notifHandler := handler.NewNotificationHandler(notifSvc)

	// Wire queries
	queryRepo := repository.NewQueryRepository(config.DB)
	querySvc := service.NewQueryService(queryRepo)
	queryHandler := handler.NewQueryHandler(querySvc)

	// Wire leave
	leaveRepo := repository.NewLeaveRepository(config.DB)
	leaveSvc := service.NewLeaveService(leaveRepo)
	leaveHandler := handler.NewLeaveHandler(leaveSvc)

	// Wire payroll
	payrollRepo := repository.NewPayrollRepository(config.DB)
	payrollSvc := service.NewPayrollService(payrollRepo)
	payrollHandler := handler.NewPayrollHandler(payrollSvc)

	// Wire documents
	docRepo := repository.NewDocumentRepository(config.DB)
	docSvc := service.NewDocumentService(docRepo)
	docHandler := handler.NewDocumentHandler(docSvc)

	// Wire employee
	empRepo := repository.NewEmployeeRepository(config.DB)
	employeeUserRepo := repository.NewUserRepository(config.DB)
	empSvc := service.NewEmployeeService(empRepo, employeeUserRepo)
	empHandler := handler.NewEmployeeHandler(empSvc)

	// Wire attendance
	attRepo := repository.NewAttendanceRepository(config.DB)
	attSvc := service.NewAttendanceService(attRepo)
	attHandler := handler.NewAttendanceHandler(attSvc)

	api := router.Group("/api")
	routes.RegisterAuthRoutes(api, authHandler)
	routes.RegisterProtectedRoutes(api, userHandler, notifHandler, queryHandler, leaveHandler, payrollHandler, docHandler, empHandler, attHandler)

	port := config.GetEnv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Println("✅ Server started at :" + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
