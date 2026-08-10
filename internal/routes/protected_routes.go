package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/handler"
	"github.com/smustafa-tech/hrms-backend/internal/middleware"
)

func RegisterProtectedRoutes(api *gin.RouterGroup, empHandler *handler.EmployeeHandler, attHandler *handler.AttendanceHandler) {
	auth := middleware.AuthMiddleware()

	// Users
	users := api.Group("/users", auth)
	{
		users.GET("/me", handler.GetMe)
	}

	// Notifications
	notifications := api.Group("/notifications", auth)
	{
		notifications.GET("/", handler.GetNotifications)
		notifications.GET("/unread-count", handler.GetNotificationsUnreadCount)
		notifications.PATCH("/:id/read", handler.MarkNotificationRead)
		notifications.PATCH("/read-all", handler.MarkAllNotificationsRead)
	}

	// Queries
	queries := api.Group("/queries", auth)
	{
		queries.GET("/unread-count", handler.GetQueriesUnreadCount)
		queries.GET("/my", handler.GetMyQueries)
		queries.GET("/all", handler.GetAllQueries)
		queries.POST("/", handler.SubmitQuery)
		queries.POST("/:id/reply", handler.ReplyToQuery)
		queries.PATCH("/:id/close", handler.CloseQuery)
		queries.DELETE("/:id", handler.DeleteQuery)
	}

	// Employees
	employees := api.Group("/employee", auth)
	{
		employees.GET("/employees-data", empHandler.GetEmployeesData)
		employees.PUT("/update-employee-data/:emp_id", empHandler.UpdateEmployee)
		employees.DELETE("/delete-employee-data/:emp_id", empHandler.DeleteEmployee)
	}

	// Add employee (separate prefix used by frontend)
	addEmp := api.Group("/add-employee", auth)
	{
		addEmp.POST("/register", empHandler.CreateEmployee)
	}

	// Attendance
	attendance := api.Group("/attendance", auth)
	{
		attendance.GET("/all-employee-attendance", attHandler.GetAllAttendance)
		attendance.GET("/own-attendance", attHandler.GetOwnAttendance)
		attendance.GET("/employee-data", attHandler.GetEmployeeData)
		attendance.POST("/mark", attHandler.MarkAttendance)
		attendance.POST("/admin-mark", attHandler.AdminMarkAttendance)
		attendance.POST("/start-break", attHandler.StartBreak)
		attendance.POST("/end-break", attHandler.EndBreak)
	}
}
