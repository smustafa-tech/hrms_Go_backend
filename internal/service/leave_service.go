package service

import (
	"errors"
	"time"

	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type LeaveService struct {
	repo *repository.LeaveRepository
}

func NewLeaveService(repo *repository.LeaveRepository) *LeaveService {
	return &LeaveService{repo: repo}
}

func (s *LeaveService) GetAllLeaves() ([]models.Leave, error) {
	return s.repo.FindAll()
}

func (s *LeaveService) GetMyLeaves(userID string) ([]models.Leave, error) {
	return s.repo.FindByUserID(userID)
}

func (s *LeaveService) ApplyLeave(userID string, req dto.ApplyLeaveRequest) (*models.Leave, error) {
	if req.LeaveType == "" || req.StartDate == "" || req.EndDate == "" {
		return nil, errors.New("missing leave type or dates")
	}

	start, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date")
	}
	end, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, errors.New("invalid end date")
	}
	if end.Before(start) {
		return nil, errors.New("end date must not be before start date")
	}

	totalDays := int(end.Sub(start).Hours()/24) + 1
	if req.TotalDays > 0 {
		totalDays = req.TotalDays
	}

	leave := &models.Leave{
		UserID:    userID,
		LeaveType: req.LeaveType,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		TotalDays: totalDays,
		Reason:    req.Reason,
		Status:    "PENDING",
	}

	if err := s.repo.Create(leave); err != nil {
		return nil, err
	}
	return leave, nil
}

func (s *LeaveService) UpdateLeaveStatus(userID string, req dto.UpdateLeaveRequest) (*models.Leave, error) {
	if _, err := s.repo.FindByID(req.LeaveID); err != nil {
		return nil, errors.New("leave not found")
	}

	updates := map[string]interface{}{
		"status": req.Status,
	}
	if req.Status != "" {
		updates["approved_by"] = userID
	}

	updated, err := s.repo.UpdateStatus(req.LeaveID, updates)
	if err != nil {
		return nil, err
	}

	if req.Status == "APPROVED" {
		if err := s.deductBalance(updated); err != nil {
			return nil, err
		}
	}

	return updated, nil
}

func (s *LeaveService) deductBalance(leave *models.Leave) error {
	year, err := time.Parse("2006-01-02", leave.StartDate)
	if err != nil {
		return nil
	}

	balance, err := s.repo.FindBalance(leave.UserID, year.Year())
	if err != nil {
		balance = &models.LeaveBalance{
			EmployeeID:   leave.UserID,
			Year:         year.Year(),
		}
		if err := s.repo.CreateBalance(balance); err != nil {
			return err
		}
	}

	switch leave.LeaveType {
	case "casual":
		balance.UsedCasual += leave.TotalDays
	case "sick":
		balance.UsedSick += leave.TotalDays
	case "earned":
		balance.UsedEarned += leave.TotalDays
	case "optional":
		balance.UsedOptional += leave.TotalDays
	default:
		balance.UsedCasual += leave.TotalDays
	}

	return s.repo.SaveBalance(balance)
}

func (s *LeaveService) GetAllLeaveBalances() ([]models.LeaveBalance, error) {
	return s.repo.FindBalances()
}

func (s *LeaveService) GetMyLeaveBalance(userID string) (*models.LeaveBalance, error) {
	year := time.Now().Year()
	balance, err := s.repo.FindBalance(userID, year)
	if err != nil {
		balance = &models.LeaveBalance{
			EmployeeID:  userID,
			Year:        year,
		}
		if err := s.repo.CreateBalance(balance); err != nil {
			return nil, err
		}
	}
	return balance, nil
}
