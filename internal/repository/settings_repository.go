package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type SettingsRepository struct {
	db *gorm.DB
}

func NewSettingsRepository(db *gorm.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (r *SettingsRepository) FindByUserID(userID string) (*models.SystemSettings, error) {
	var settings models.SystemSettings
	err := r.db.Where("user_id = ?", userID).First(&settings).Error
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func (r *SettingsRepository) Create(settings *models.SystemSettings) error {
	return r.db.Create(settings).Error
}

func (r *SettingsRepository) Update(userID string, updates map[string]interface{}) (*models.SystemSettings, error) {
	if err := r.db.Model(&models.SystemSettings{}).Where("user_id = ?", userID).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.FindByUserID(userID)
}
