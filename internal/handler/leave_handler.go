package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type LeaveHandler struct {
	svc *service.LeaveService
}

func NewLeaveHandler(svc *service.LeaveService) *LeaveHandler {
	return &LeaveHandler{svc: svc}
}

func (h *LeaveHandler) GetAllLeaves(c *gin.Context) {
	leaves, err := h.svc.GetAllLeaves()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"allLeaves": leaves})
}

func (h *LeaveHandler) GetMyLeaves(c *gin.Context) {
	userID, _ := c.Get("userID")
	leaves, err := h.svc.GetMyLeaves(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"leaves": leaves})
}

func (h *LeaveHandler) ApplyLeave(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.ApplyLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	leave, err := h.svc.ApplyLeave(userID.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"leave": leave})
}

func (h *LeaveHandler) UpdateLeave(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.UpdateLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	leave, err := h.svc.UpdateLeaveStatus(userID.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"leave": leave})
}

func (h *LeaveHandler) GetAllLeaveBalances(c *gin.Context) {
	balances, err := h.svc.GetAllLeaveBalances()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balances": balances})
}

func (h *LeaveHandler) GetMyLeaveBalance(c *gin.Context) {
	userID, _ := c.Get("userID")
	balance, err := h.svc.GetMyLeaveBalance(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
