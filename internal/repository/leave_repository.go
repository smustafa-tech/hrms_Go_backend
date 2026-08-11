package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type LeaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) *LeaveRepository {
	return &LeaveRepository{db: db}
}

func (r *LeaveRepository) FindAll() ([]models.Leave, error) {
	var records []models.Leave
	err := r.db.Order("created_at desc").Find(&records).Error
	return records, err
}

func (r *LeaveRepository) FindByUserID(userID string) ([]models.Leave, error) {
	var records []models.Leave
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&records).Error
	return records, err
}

func (r *LeaveRepository) Create(record *models.Leave) error {
	return r.db.Create(record).Error
}

func (r *LeaveRepository) UpdateStatus(id string, updates map[string]interface{}) (*models.Leave, error) {
	if err := r.db.Model(&models.Leave{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindByID(id)
}

func (r *LeaveRepository) FindByID(id string) (*models.Leave, error) {
	var record models.Leave
	err := r.db.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *LeaveRepository) FindBalances() ([]models.LeaveBalance, error) {
	var balances []models.LeaveBalance
	err := r.db.Order("year desc").Find(&balances).Error
	return balances, err
}

func (r *LeaveRepository) FindBalance(employeeID string, year int) (*models.LeaveBalance, error) {
	var balance models.LeaveBalance
	err := r.db.Where("employee_id = ? AND year = ?", employeeID, year).First(&balance).Error
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (r *LeaveRepository) CreateBalance(balance *models.LeaveBalance) error {
	return r.db.Create(balance).Error
}

func (r *LeaveRepository) SaveBalance(balance *models.LeaveBalance) error {
	return r.db.Save(balance).Error
}
