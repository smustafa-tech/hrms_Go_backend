package service

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(repo *repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

func (s *CompanyService) GetCompany() (*models.Company, error) {
	company, err := s.repo.FindFirst()
	if err != nil {
		return &models.Company{}, nil
	}
	return company, nil
}

func (s *CompanyService) UpdateCompany(req *models.Company) (*models.Company, error) {
	company, err := s.repo.FindFirst()
	if err != nil {
		if err := s.repo.Create(req); err != nil {
			return nil, err
		}
		return req, nil
	}
	return s.repo.Update(company.ID.String(), map[string]interface{}{
		"name":    req.Name,
		"email":   req.Email,
		"phone":   req.Phone,
		"address": req.Address,
	})
}
