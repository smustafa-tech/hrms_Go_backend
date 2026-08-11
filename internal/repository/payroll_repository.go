package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type PayrollRepository struct {
	db *gorm.DB
}

func NewPayrollRepository(db *gorm.DB) *PayrollRepository {
	return &PayrollRepository{db: db}
}

func (r *PayrollRepository) FindAll() ([]models.Payroll, error) {
	var payrolls []models.Payroll
	err := r.db.Order("year desc, month desc").Find(&payrolls).Error
	return payrolls, err
}

func (r *PayrollRepository) FindByUserID(userID string) ([]models.Payroll, error) {
	var payrolls []models.Payroll
	err := r.db.Where("user_id = ?", userID).Order("year desc, month desc").Find(&payrolls).Error
	return payrolls, err
}

func (r *PayrollRepository) FindByID(id string) (*models.Payroll, error) {
	var payroll models.Payroll
	err := r.db.Where("id = ?", id).First(&payroll).Error
	if err != nil {
		return nil, err
	}
	return &payroll, nil
}

func (r *PayrollRepository) Create(payroll *models.Payroll) error {
	return r.db.Create(payroll).Error
}

func (r *PayrollRepository) Update(id string, updates map[string]interface{}) (*models.Payroll, error) {
	if err := r.db.Model(&models.Payroll{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindByID(id)
}
