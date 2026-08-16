package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct{}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) EmployeeDashboardStats(
	c *gin.Context,
) {

	userID, _ := c.Get("userID")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(
		http.StatusOK,
		gin.H{

			"success": true,

			"message":
			"Employee dashboard data fetched successfully",

			"data": gin.H{

				"userID": userID,

				"email": email,

				"role": role,
			},
		},
	)
}

func (h *DashboardHandler) ManagerDashboardStats(
	c *gin.Context,
) {

	userID, _ := c.Get("userID")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message":
			"Manager dashboard data fetched successfully",
			"data": gin.H{
				"userID": userID,
				"email":  email,
				"role":   role,
			},
		},
	)
}

func (h *DashboardHandler) HRDashboardStats(
	c *gin.Context,
) {

	userID, _ := c.Get("userID")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message":
			"HR dashboard data fetched successfully",
			"data": gin.H{
				"userID": userID,
				"email":  email,
				"role":   role,
			},
		},
	)
}

func (h *DashboardHandler) AdminDashboardStats(
	c *gin.Context,
) {

	userID, _ := c.Get("userID")
	email, _ := c.Get("email")
	role, _ := c.Get("role")

	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"message":
			"Admin dashboard data fetched successfully",
			"data": gin.H{
				"userID": userID,
				"email":  email,
				"role":   role,
			},
		},
	)
}