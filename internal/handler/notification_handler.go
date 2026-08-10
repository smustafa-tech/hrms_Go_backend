package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotificationsUnreadCount(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"unreadCount": 0, "count": 0})
}

func GetNotifications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"notifications": []interface{}{}})
}

func MarkNotificationRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "marked as read"})
}

func MarkAllNotificationsRead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all marked as read"})
}
