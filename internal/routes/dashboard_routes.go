package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/smustafa-tech/hrms-backend/internal/handler"
	"github.com/smustafa-tech/hrms-backend/internal/middleware"
)

func RegisterDashboardRoutes(
	api *gin.RouterGroup,
	dashboardHandler *handler.DashboardHandler,
) {

	dashboard :=
		api.Group("/dashboard")

	// Employee dashboard
	dashboard.GET(
		"/dashboard-stats",

		middleware.AuthMiddleware(),

		middleware.RequireRoles(
			"employee",
		),

		dashboardHandler.EmployeeDashboardStats,
	)

	// Manager dashboard
	dashboard.GET(
		"/manager-dashboard-stats",

		middleware.AuthMiddleware(),

		middleware.RequireRoles(
			"manager",
		),

		dashboardHandler.ManagerDashboardStats,
	)

	// HR dashboard
	dashboard.GET(
		"/hr-dashboard-stats",

		middleware.AuthMiddleware(),

		middleware.RequireRoles(
			"hr",
		),

		dashboardHandler.HRDashboardStats,
	)

	// Admin dashboard
	dashboard.GET(
		"/admin-dashboard-stats",

		middleware.AuthMiddleware(),

		middleware.RequireRoles(
			"admin",
		),

		dashboardHandler.AdminDashboardStats,
	)
}