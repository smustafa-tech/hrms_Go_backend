package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/config"
	"github.com/smustafa-tech/hrms-backend/internal/models"
)

func GetMe(c *gin.Context) {
	email, _ := c.Get("email")

	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}
