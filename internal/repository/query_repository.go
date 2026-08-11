package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type QueryRepository struct {
	db *gorm.DB
}

func NewQueryRepository(db *gorm.DB) *QueryRepository {
	return &QueryRepository{db: db}
}

func (r *QueryRepository) FindByUserID(userID string) ([]models.Query, error) {
	var queries []models.Query
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&queries).Error
	return queries, err
}

func (r *QueryRepository) FindAll() ([]models.Query, error) {
	var queries []models.Query
	err := r.db.Order("created_at desc").Find(&queries).Error
	return queries, err
}

func (r *QueryRepository) FindByID(id string) (*models.Query, error) {
	var query models.Query
	err := r.db.Where("id = ?", id).First(&query).Error
	if err != nil {
		return nil, err
	}
	return &query, nil
}

func (r *QueryRepository) Create(query *models.Query) error {
	return r.db.Create(query).Error
}

func (r *QueryRepository) CreateReply(reply *models.QueryReply) error {
	return r.db.Create(reply).Error
}

func (r *QueryRepository) FindRepliesForQueryIDs(queryIDs []string) ([]models.QueryReply, error) {
	var replies []models.QueryReply
	err := r.db.Where("query_id IN ?", queryIDs).Order("created_at asc").Find(&replies).Error
	return replies, err
}

func (r *QueryRepository) CloseQuery(id string) error {
	return r.db.Model(&models.Query{}).Where("id = ?", id).Updates(map[string]interface{}{"status": "closed"}).Error
}

func (r *QueryRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.Query{}).Error
}

func (r *QueryRepository) CountOpenByUser(userID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Query{}).Where("user_id = ? AND status != ?", userID, "closed").Count(&count).Error
	return count, err
}
