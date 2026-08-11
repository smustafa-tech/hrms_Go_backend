package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/service"
)

type UserHandler struct {
	userSvc     *service.UserService
	companySvc  *service.CompanyService
	settingsSvc *service.SettingsService
}

func NewUserHandler(userSvc *service.UserService, companySvc *service.CompanyService, settingsSvc *service.SettingsService) *UserHandler {
	return &UserHandler{userSvc: userSvc, companySvc: companySvc, settingsSvc: settingsSvc}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID, _ := c.Get("userID")
	user, err := h.userSvc.GetProfile(userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": user})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.userSvc.UpdateProfile(userID.(string), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated", "profile": user})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := h.userSvc.ChangePassword(userID.(string), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}

func (h *UserHandler) UploadProfilePhoto(c *gin.Context) {
	userID, _ := c.Get("userID")
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "profile photo is required"})
		return
	}

	uploadDir := filepath.Join("uploads", "profiles")
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create upload directory"})
		return
	}

	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, fileName)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save photo"})
		return
	}

	user, err := h.userSvc.UploadProfilePhoto(userID.(string), filePath, "/uploads/profiles/"+fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile photo uploaded", "profile": user})
}

func (h *UserHandler) GetNotificationSettings(c *gin.Context) {
	userID, _ := c.Get("userID")
	settings, err := h.settingsSvc.GetUserSettings(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *UserHandler) UpdateNotificationSettings(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.NotificationSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if req.EmailNotifications != nil {
		updates["email_notifications"] = *req.EmailNotifications
	}
	if req.PushNotifications != nil {
		updates["push_notifications"] = *req.PushNotifications
	}
	if req.LeaveRequests != nil {
		updates["leave_requests"] = *req.LeaveRequests
	}
	if req.AttendanceAlerts != nil {
		updates["attendance_alerts"] = *req.AttendanceAlerts
	}

	settings, err := h.settingsSvc.UpdateUserSettings(userID.(string), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *UserHandler) GetSecuritySettings(c *gin.Context) {
	userID, _ := c.Get("userID")
	settings, err := h.settingsSvc.GetUserSettings(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *UserHandler) UpdateSecuritySettings(c *gin.Context) {
	userID, _ := c.Get("userID")
	var req dto.SecuritySettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updates := map[string]interface{}{}
	if req.TwoFactorAuth != nil {
		updates["two_factor_auth"] = *req.TwoFactorAuth
	}
	if req.SessionTimeout != "" {
		updates["session_timeout"] = req.SessionTimeout
	}

	settings, err := h.settingsSvc.UpdateUserSettings(userID.(string), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

func (h *UserHandler) GetCompany(c *gin.Context) {
	company, err := h.companySvc.GetCompany()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *UserHandler) UpdateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	updated, err := h.companySvc.UpdateCompany(&company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}
