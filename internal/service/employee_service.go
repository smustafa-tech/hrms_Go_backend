package service

import (
	"errors"

	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type EmployeeService struct {
	repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func salaryVal(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}

func (s *EmployeeService) GetAllEmployees() (map[string]interface{}, error) {
	employees, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	active := []models.Employee{}
	inactive := []models.Employee{}
	for _, e := range employees {
		if e.Status == "active" {
			active = append(active, e)
		} else {
			inactive = append(inactive, e)
		}
	}

	deptCounts := s.repo.DepartmentCounts()

	return map[string]interface{}{
		"totalEmployees":       employees,
		"totalEmployeesCount":  len(employees),
		"activeEmployeesCount": len(active),
		"departmentWiseCounts": deptCounts,
	}, nil
}

func (s *EmployeeService) CreateEmployee(req dto.CreateEmployeeRequest) (*models.Employee, error) {
	if s.repo.EmpIDExists(req.EmpID) {
		return nil, errors.New("employee ID already exists")
	}
	if s.repo.EmailExists(req.Email) {
		return nil, errors.New("email already registered")
	}

	status := req.Status
	if status == "" {
		status = "active"
	}
	role := req.Role
	if role == "" {
		role = "employee"
	}
	workMode := req.WorkMode
	if workMode == "" {
		workMode = "office"
	}

	emp := &models.Employee{
		EmpID:         req.EmpID,
		FirstName:     req.FirstName,
		MiddleName:    req.MiddleName,
		LastName:      req.LastName,
		Email:         req.Email,
		Phone:         req.Phone,
		AdharCard:     req.AdharCard,
		Designation:   req.Designation,
		Role:          role,
		Department:    req.Department,
		DateOfJoining: req.DateOfJoining,
		Status:        status,
		Salary:        salaryVal(req.Salary),
		WorkMode:      workMode,
		MgrID:         req.MgrID,
		HrID:          req.HrID,
	}

	if err := s.repo.Create(emp); err != nil {
		return nil, errors.New("failed to create employee")
	}
	return emp, nil
}

func (s *EmployeeService) UpdateEmployee(empID string, req dto.UpdateEmployeeRequest) (*models.Employee, error) {
	updates := map[string]interface{}{}

	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.MiddleName != "" {
		updates["middle_name"] = req.MiddleName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.AdharCard != "" {
		updates["adhar_card"] = req.AdharCard
	}
	if req.Designation != "" {
		updates["designation"] = req.Designation
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.Department != "" {
		updates["department"] = req.Department
	}
	if req.DateOfJoining != "" {
		updates["date_of_joining"] = req.DateOfJoining
	}
	if req.Status != "" {
		updates["status"] = req.Status
	}
	if req.Salary != nil && *req.Salary > 0 {
		updates["salary"] = *req.Salary
	}
	if req.WorkMode != "" {
		updates["work_mode"] = req.WorkMode
	}
	if req.MgrID != "" {
		updates["mgr_id"] = req.MgrID
	}
	if req.HrID != "" {
		updates["hr_id"] = req.HrID
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	emp, err := s.repo.Update(empID, updates)
	if err != nil {
		return nil, errors.New("employee not found or update failed")
	}
	return emp, nil
}

func (s *EmployeeService) DeleteEmployee(empID string) error {
	if !s.repo.EmpIDExists(empID) {
		return errors.New("employee not found")
	}
	return s.repo.SoftDelete(empID)
}
