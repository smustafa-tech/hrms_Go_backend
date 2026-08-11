package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/handler"
	"github.com/smustafa-tech/hrms-backend/internal/middleware"
)

func RegisterProtectedRoutes(
	api *gin.RouterGroup,
	userHandler *handler.UserHandler,
	notificationHandler *handler.NotificationHandler,
	queryHandler *handler.QueryHandler,
	leaveHandler *handler.LeaveHandler,
	payrollHandler *handler.PayrollHandler,
	documentHandler *handler.DocumentHandler,
	empHandler *handler.EmployeeHandler,
	attHandler *handler.AttendanceHandler,
) {
	auth := middleware.AuthMiddleware()

	// Users
	users := api.Group("/users", auth)
	{
		users.GET("/me", userHandler.GetMe)
		users.PUT("/update", userHandler.UpdateProfile)
		users.PUT("/change-password", userHandler.ChangePassword)
		users.PUT("/profile-photo", userHandler.UploadProfilePhoto)
		users.GET("/notifications", userHandler.GetNotificationSettings)
		users.PUT("/notifications", userHandler.UpdateNotificationSettings)
		users.GET("/security-settings", userHandler.GetSecuritySettings)
		users.PUT("/security-settings", userHandler.UpdateSecuritySettings)
	}

	// Company
	company := api.Group("/company", auth)
	{
		company.GET("/me", userHandler.GetCompany)
		company.PUT("/update", userHandler.UpdateCompany)
	}

	// Notifications
	notifications := api.Group("/notifications", auth)
	{
		notifications.GET("/", notificationHandler.GetNotifications)
		notifications.GET("/unread-count", notificationHandler.GetNotificationsUnreadCount)
		notifications.PATCH("/:id/read", notificationHandler.MarkNotificationRead)
		notifications.PATCH("/read-all", notificationHandler.MarkAllNotificationsRead)
	}

	// Queries
	queries := api.Group("/queries", auth)
	{
		queries.GET("/unread-count", queryHandler.GetQueriesUnreadCount)
		queries.GET("/my", queryHandler.GetMyQueries)
		queries.GET("/all", queryHandler.GetAllQueries)
		queries.POST("/", queryHandler.SubmitQuery)
		queries.POST("/:id/reply", queryHandler.ReplyToQuery)
		queries.PATCH("/:id/close", queryHandler.CloseQuery)
		queries.DELETE("/:id", queryHandler.DeleteQuery)
	}

	// Employees
	employees := api.Group("/employee", auth, middleware.RequireRoles("admin", "hr", "manager"))
	{
		employees.GET("/employees-data", empHandler.GetEmployeesData)
		employees.PUT("/update-employee-data/:emp_id", empHandler.UpdateEmployee)
		employees.DELETE("/delete-employee-data/:emp_id", empHandler.DeleteEmployee)
	}

	// Add employee
	addEmp := api.Group("/add-employee", auth, middleware.RequireRoles("admin", "hr", "manager"))
	{
		addEmp.POST("/register", empHandler.CreateEmployee)
	}

	// Attendance
	attendance := api.Group("/attendance", auth)
	{
		attendance.GET("/all-employee-attendance", middleware.RequireRoles("admin", "hr", "manager"), attHandler.GetAllAttendance)
		attendance.GET("/own-attendance", attHandler.GetOwnAttendance)
		attendance.GET("/employee-data", middleware.RequireRoles("admin", "hr", "manager"), attHandler.GetEmployeeData)
		attendance.POST("/mark", attHandler.MarkAttendance)
		attendance.POST("/admin-mark", middleware.RequireRoles("admin", "hr"), attHandler.AdminMarkAttendance)
		attendance.POST("/start-break", attHandler.StartBreak)
		attendance.POST("/end-break", attHandler.EndBreak)
	}

	// Leave
	leave := api.Group("/leave", auth)
	{
		leave.GET("/all-leaves", leaveHandler.GetAllLeaves)
		leave.GET("/my-leaves", leaveHandler.GetMyLeaves)
		leave.POST("/apply-leave", leaveHandler.ApplyLeave)
		leave.PUT("/update-leave", leaveHandler.UpdateLeave)
		leave.GET("/all-leaveBalance", leaveHandler.GetAllLeaveBalances)
		leave.GET("/my-leaveBalance", leaveHandler.GetMyLeaveBalance)
	}

	// Payroll
	payroll := api.Group("/payroll", auth)
	{
		payroll.GET("/list", payrollHandler.GetAllPayrolls)
		payroll.GET("/my-payrolls", payrollHandler.GetMyPayrolls)
		payroll.GET("/employees", payrollHandler.GetEmployeesForPayroll)
		payroll.POST("/create", payrollHandler.CreatePayroll)
		payroll.PUT("/update/:id", payrollHandler.UpdatePayroll)
		payroll.PUT("/status/:id", payrollHandler.UpdatePayrollStatus)
		payroll.GET("/payslip/:id/pdf", payrollHandler.GetPayslipPDF)
	}

	// Documents
	docs := api.Group("/document", auth)
	{
		docs.GET("/me", documentHandler.GetMyDocuments)
		docs.GET("/employee/all", documentHandler.GetEmployeeDocuments)
		docs.POST("/upload", documentHandler.UploadDocument)
		docs.GET("/download/:id", documentHandler.DownloadDocument)
		docs.DELETE("/:id", documentHandler.DeleteDocument)
	}
}
