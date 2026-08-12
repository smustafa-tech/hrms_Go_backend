package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"

	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeService struct {
	repo     *repository.EmployeeRepository
	userRepo *repository.UserRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository, userRepo *repository.UserRepository) *EmployeeService {
	return &EmployeeService{repo: repo, userRepo: userRepo}
}

func salaryVal(s *float64) float64 {
	if s == nil {
		return 0
	}
	return *s
}

func generateTemporaryPassword() string {
	buf := make([]byte, 8)
	if _, err := rand.Read(buf); err != nil {
		return "Temp@1234"
	}
	return "Temp@" + hex.EncodeToString(buf)[:8]
}

func normalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

func normalizeRole(role string) string {
	return strings.ToLower(strings.TrimSpace(role))
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

func (s *EmployeeService) CreateEmployee(req dto.CreateEmployeeRequest, slug string) (*models.Employee, string, error) {
	firstName := strings.TrimSpace(req.FirstName)
	middleName := strings.TrimSpace(req.MiddleName)
	lastName := strings.TrimSpace(req.LastName)
	email := normalizeEmail(req.Email)
	phone := strings.TrimSpace(req.Phone)
	adharCard := strings.TrimSpace(req.AdharCard)
	designation := strings.TrimSpace(req.Designation)
	role := normalizeRole(req.Role)
	if role == "" {
		role = "employee"
	}
	department := strings.TrimSpace(req.Department)
	workMode := strings.TrimSpace(req.WorkMode)
	if workMode == "" {
		workMode = "office"
	}

	if firstName == "" || lastName == "" || email == "" || designation == "" || department == "" || req.EmpID == "" || strings.TrimSpace(req.TemporaryPassword) == "" {
		return nil, "", errors.New("all required employee fields and temporary password must be provided")
	}
	if !regexp.MustCompile(`^\+91\s?[6-9]\d{9}$`).MatchString(phone) {
		return nil, "", errors.New("phone number must be in format +91XXXXXXXXXX")
	}
	if adharCard == "" {
		return nil, "", errors.New("aadhaar number is required")
	}
	if s.repo.EmpIDExists(req.EmpID) {
		return nil, "", errors.New("employee ID already exists")
	}
	if s.repo.EmailExists(email) {
		return nil, "", errors.New("email already registered")
	}
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return nil, "", errors.New("employee login already exists for this email")
	}

	status := req.Status
	if status == "" {
		status = "active"
	}

	emp := &models.Employee{
		EmpID:         req.EmpID,
		FirstName:     firstName,
		MiddleName:    middleName,
		LastName:      lastName,
		Email:         email,
		Phone:         phone,
		AdharCard:     adharCard,
		Designation:   designation,
		Role:          role,
		Department:    department,
		DateOfJoining: req.DateOfJoining,
		Status:        status,
		Salary:        salaryVal(req.Salary),
		WorkMode:      workMode,
		MgrID:         req.MgrID,
		HrID:          req.HrID,
	}

	if err := s.repo.Create(emp); err != nil {
		return nil, "", errors.New("failed to create employee")
	}

	tempPassword := strings.TrimSpace(req.TemporaryPassword)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tempPassword), bcrypt.DefaultCost)
	if err != nil {
		_ = s.repo.DeleteByEmpID(req.EmpID)
		return nil, "", errors.New("failed to generate employee login credentials")
	}

	if slug == "" {
		slug = "default"
	}
	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
		Phone:     phone,
		Slug:      slug,
		Role:      role,
	}
	if err := s.userRepo.Create(user); err != nil {
		_ = s.repo.DeleteByEmpID(req.EmpID)
		return nil, "", errors.New("failed to create employee login credentials")
	}

	return emp, tempPassword, nil
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
