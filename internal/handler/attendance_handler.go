package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type AttendanceHandler struct {
	svc *service.AttendanceService
}

func NewAttendanceHandler(svc *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{svc: svc}
}

// POST /attendance/mark — employee check-in / check-out
func (h *AttendanceHandler) MarkAttendance(c *gin.Context) {
	var req dto.MarkAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record, err := h.svc.MarkAttendance(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked", "attendance": record})
}

// POST /attendance/admin-mark — admin/hr mark for any employee
func (h *AttendanceHandler) AdminMarkAttendance(c *gin.Context) {
	var req dto.AdminMarkAttendanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record, err := h.svc.AdminMarkAttendance(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Attendance marked", "attendance": record})
}

// POST /attendance/start-break
func (h *AttendanceHandler) StartBreak(c *gin.Context) {
	var req dto.BreakRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record, err := h.svc.StartBreak(req.AttendanceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Break started", "attendance": record})
}

// POST /attendance/end-break
func (h *AttendanceHandler) EndBreak(c *gin.Context) {
	var req dto.BreakRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	record, err := h.svc.EndBreak(req.AttendanceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Break ended", "attendance": record})
}

// GET /attendance/own-attendance — employee's own records
func (h *AttendanceHandler) GetOwnAttendance(c *gin.Context) {
	userID, _ := c.Get("userID")

	data, err := h.svc.GetOwnAttendance(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GET /attendance/all-employee-attendance — admin/hr
func (h *AttendanceHandler) GetAllAttendance(c *gin.Context) {
	data, err := h.svc.GetAllAttendance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

// GET /attendance/employee-data
func (h *AttendanceHandler) GetEmployeeData(c *gin.Context) {
	data, err := h.svc.GetAllAttendance()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
