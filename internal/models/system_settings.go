package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SystemSettings struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID             string    `gorm:"not null;index" json:"userId"`
	Theme              string    `gorm:"default:'light'" json:"theme"`
	Language           string    `gorm:"default:'en'" json:"language"`
	EmailNotifications bool      `gorm:"default:true" json:"emailNotifications"`
	PushNotifications  bool      `gorm:"default:false" json:"pushNotifications"`
	LeaveRequests      bool      `gorm:"default:true" json:"leaveRequests"`
	AttendanceAlerts   bool      `gorm:"default:true" json:"attendanceAlerts"`
	TwoFactorAuth      bool      `gorm:"default:false" json:"twoFactorAuth"`
	SessionTimeout     string    `gorm:"default:'60'" json:"sessionTimeout"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

func (s *SystemSettings) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New()
	return nil
}
