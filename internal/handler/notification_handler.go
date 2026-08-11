package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type NotificationHandler struct {
	svc *service.NotificationService
}

func NewNotificationHandler(svc *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{svc: svc}
}

func (h *NotificationHandler) GetNotificationsUnreadCount(c *gin.Context) {
	userID, _ := c.Get("userID")
	count, err := h.svc.GetUnreadCount(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"unreadCount": count})
}

func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userID, _ := c.Get("userID")
	notifications, err := h.svc.GetMyNotifications(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"notifications": notifications})
}

func (h *NotificationHandler) MarkNotificationRead(c *gin.Context) {
	id := c.Param("id")
	if err := h.svc.MarkAsRead(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "notification marked as read"})
}

func (h *NotificationHandler) MarkAllNotificationsRead(c *gin.Context) {
	userID, _ := c.Get("userID")
	if err := h.svc.MarkAllAsRead(userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "all notifications marked as read"})
}
