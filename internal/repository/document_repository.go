package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (r *DocumentRepository) FindByUserID(userID string) ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) FindAll() ([]models.Document, error) {
	var documents []models.Document
	err := r.db.Order("created_at desc").Find(&documents).Error
	return documents, err
}

func (r *DocumentRepository) FindByID(id string) (*models.Document, error) {
	var doc models.Document
	err := r.db.Where("id = ?", id).First(&doc).Error
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *DocumentRepository) Create(doc *models.Document) error {
	return r.db.Create(doc).Error
}

func (r *DocumentRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Document{}).Error
}
