package service

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type NotificationService struct {
	repo *repository.NotificationRepository
}

func NewNotificationService(repo *repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) GetMyNotifications(receiverID string) ([]models.Notification, error) {
	return s.repo.FindByReceiver(receiverID)
}

func (s *NotificationService) GetUnreadCount(receiverID string) (int64, error) {
	return s.repo.CountUnread(receiverID)
}

func (s *NotificationService) MarkAsRead(id string) error {
	return s.repo.MarkRead(id)
}

func (s *NotificationService) MarkAllAsRead(receiverID string) error {
	return s.repo.MarkAllRead(receiverID)
}
