package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) FindByReceiver(receiverID string) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("receiver_id = ?", receiverID).Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) CountUnread(receiverID string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Notification{}).Where("receiver_id = ? AND is_read = false", receiverID).Count(&count).Error
	return count, err
}

func (r *NotificationRepository) MarkRead(id string) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", id).Update("is_read", true).Error
}

func (r *NotificationRepository) MarkAllRead(receiverID string) error {
	return r.db.Model(&models.Notification{}).Where("receiver_id = ? AND is_read = false", receiverID).Update("is_read", true).Error
}
