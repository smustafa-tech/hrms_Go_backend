package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) FindFirst() (*models.Company, error) {
	var company models.Company
	err := r.db.First(&company).Error
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.db.Create(company).Error
}

func (r *CompanyRepository) Update(id string, updates map[string]interface{}) (*models.Company, error) {
	if err := r.db.Model(&models.Company{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindFirst()
}
