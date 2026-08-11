package service

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type SettingsService struct {
	repo *repository.SettingsRepository
}

func NewSettingsService(repo *repository.SettingsRepository) *SettingsService {
	return &SettingsService{repo: repo}
}

func (s *SettingsService) GetUserSettings(userID string) (*models.SystemSettings, error) {
	settings, err := s.repo.FindByUserID(userID)
	if err != nil {
		settings = &models.SystemSettings{UserID: userID}
		if err := s.repo.Create(settings); err != nil {
			return nil, err
		}
		return settings, nil
	}
	return settings, nil
}

func (s *SettingsService) UpdateUserSettings(userID string, updates map[string]interface{}) (*models.SystemSettings, error) {
	settings, err := s.repo.FindByUserID(userID)
	if err != nil {
		settings = &models.SystemSettings{UserID: userID}
		if err := s.repo.Create(settings); err != nil {
			return nil, err
		}
	}
	return s.repo.Update(userID, updates)
}
