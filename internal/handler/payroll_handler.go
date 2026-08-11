package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type PayrollHandler struct {
	svc *service.PayrollService
}

func NewPayrollHandler(svc *service.PayrollService) *PayrollHandler {
	return &PayrollHandler{svc: svc}
}

func (h *PayrollHandler) GetAllPayrolls(c *gin.Context) {
	payrolls, err := h.svc.GetAllPayrolls()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payrolls": payrolls})
}

func (h *PayrollHandler) GetMyPayrolls(c *gin.Context) {
	userID, _ := c.Get("userID")
	payrolls, err := h.svc.GetMyPayrolls(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payrolls": payrolls})
}

func (h *PayrollHandler) GetEmployeesForPayroll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"employees": []interface{}{}})
}

func (h *PayrollHandler) CreatePayroll(c *gin.Context) {
	var req dto.PayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payroll, err := h.svc.CreatePayroll(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"payroll": payroll})
}

func (h *PayrollHandler) UpdatePayroll(c *gin.Context) {
	id := c.Param("id")
	var req dto.PayrollRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payroll, err := h.svc.UpdatePayroll(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payroll": payroll})
}

func (h *PayrollHandler) UpdatePayrollStatus(c *gin.Context) {
	id := c.Param("id")
	var req dto.PayrollStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	payroll, err := h.svc.UpdatePayrollStatus(id, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payroll": payroll})
}

func (h *PayrollHandler) GetPayslipPDF(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "payslip PDF generation not implemented"})
}
