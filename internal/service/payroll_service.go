package service

import (
	"errors"

	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type PayrollService struct {
	repo *repository.PayrollRepository
}

func NewPayrollService(repo *repository.PayrollRepository) *PayrollService {
	return &PayrollService{repo: repo}
}

func (s *PayrollService) GetAllPayrolls() ([]models.Payroll, error) {
	return s.repo.FindAll()
}

func (s *PayrollService) GetMyPayrolls(userID string) ([]models.Payroll, error) {
	return s.repo.FindByUserID(userID)
}

func (s *PayrollService) GetPayrollByID(id string) (*models.Payroll, error) {
	return s.repo.FindByID(id)
}

func (s *PayrollService) CreatePayroll(req dto.PayrollRequest) (*models.Payroll, error) {
	payroll := &models.Payroll{
		UserID:      req.UserID,
		Month:       req.Month,
		Year:        req.Year,
		BasicSalary: req.BasicSalary,
		HRA:         req.HRA,
		Allowances:  req.Allowances,
		Deductions:  req.Deductions,
	}
	payroll.NetSalary = payroll.BasicSalary + payroll.HRA + payroll.Allowances - payroll.Deductions
	if err := s.repo.Create(payroll); err != nil {
		return nil, err
	}
	return payroll, nil
}

func (s *PayrollService) UpdatePayroll(id string, req dto.PayrollRequest) (*models.Payroll, error) {
	updates := map[string]interface{}{
		"user_id":      req.UserID,
		"month":        req.Month,
		"year":         req.Year,
		"basic_salary": req.BasicSalary,
		"hra":          req.HRA,
		"allowances":   req.Allowances,
		"deductions":   req.Deductions,
	}
	updates["net_salary"] = req.BasicSalary + req.HRA + req.Allowances - req.Deductions
	return s.repo.Update(id, updates)
}

func (s *PayrollService) UpdatePayrollStatus(id, status string) (*models.Payroll, error) {
	if status == "" {
		return nil, errors.New("status is required")
	}
	return s.repo.Update(id, map[string]interface{}{"status": status})
}

func (s *PayrollService) GetSummary(month, year int) (map[string]interface{}, error) {
	payrolls, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var filtered []models.Payroll
	var total float64
	var count int
	for _, p := range payrolls {
		if month != 0 && p.Month != month {
			continue
		}
		if year != 0 && p.Year != year {
			continue
		}
		filtered = append(filtered, p)
		total += p.NetSalary
		count++
	}

	return map[string]interface{}{
		"summary": map[string]interface{}{
			"count":       count,
			"totalAmount": total,
		},
		"payrolls": filtered,
	}, nil
}
