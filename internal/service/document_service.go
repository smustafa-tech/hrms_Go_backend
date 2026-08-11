package service

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type DocumentService struct {
	repo *repository.DocumentRepository
}

func NewDocumentService(repo *repository.DocumentRepository) *DocumentService {
	return &DocumentService{repo: repo}
}

func (s *DocumentService) GetDocumentsForUser(userID string) ([]models.Document, error) {
	return s.repo.FindByUserID(userID)
}

func (s *DocumentService) GetAllDocuments() ([]models.Document, error) {
	return s.repo.FindAll()
}

func (s *DocumentService) CreateDocument(doc *models.Document) error {
	return s.repo.Create(doc)
}

func (s *DocumentService) FindDocumentByID(id string) (*models.Document, error) {
	return s.repo.FindByID(id)
}

func (s *DocumentService) DeleteDocument(id string) error {
	return s.repo.Delete(id)
}
