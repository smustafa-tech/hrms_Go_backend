package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/smustafa-tech/hrms-backend/internal/handler"
)

func RegisterAuthRoutes(api *gin.RouterGroup, h *handler.AuthHandler) {
	auth := api.Group("/auth")
	{
		auth.POST("/admin/register", h.Register)
		auth.POST("/user/login", h.Login)
	}
}
