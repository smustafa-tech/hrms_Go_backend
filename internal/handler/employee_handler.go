package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type EmployeeHandler struct {
	svc *service.EmployeeService
}

func NewEmployeeHandler(svc *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}

// GET /api/employee/employees-data
func (h *EmployeeHandler) GetEmployeesData(c *gin.Context) {
	data, err := h.svc.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

// POST /api/add-employee/register
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emp, err := h.svc.CreateEmployee(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Employee created successfully", "data": emp})
}

// PUT /api/employee/update-employee-data/:emp_id
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	empID := c.Param("emp_id")

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	emp, err := h.svc.UpdateEmployee(empID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully", "data": emp})
}

// DELETE /api/employee/delete-employee-data/:emp_id
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	empID := c.Param("emp_id")

	if err := h.svc.DeleteEmployee(empID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully", "success": true})
}
